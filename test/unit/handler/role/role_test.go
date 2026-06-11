package role_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"game-backend/appvalidator"
	rolehandler "game-backend/handler/role"
	"game-backend/locale"
	"game-backend/middleware"
	"game-backend/model"
	rolesvc "game-backend/service/role"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// TestMain 在所有測試前初始化 i18n 與自訂 validator
func TestMain(m *testing.M) {
	locale.Init()
	appvalidator.Init()
	m.Run()
}

// --- mock service ---

type mockRoleService struct {
	listFn    func(page, perPage int, name string) ([]model.Role, int64, error)
	getByIDFn func(id uint) (model.Role, error)
	createFn  func(name string) (model.Role, error)
	updateFn  func(id uint, name string) (model.Role, error)
	deleteFn  func(id uint) error
}

func (m *mockRoleService) List(page, perPage int, name string) ([]model.Role, int64, error) {
	return m.listFn(page, perPage, name)
}
func (m *mockRoleService) GetByID(id uint) (model.Role, error) {
	return m.getByIDFn(id)
}
func (m *mockRoleService) Create(name string) (model.Role, error) {
	return m.createFn(name)
}
func (m *mockRoleService) Update(id uint, name string) (model.Role, error) {
	return m.updateFn(id, name)
}
func (m *mockRoleService) Delete(id uint) error {
	return m.deleteFn(id)
}

// setupRouter 建立測試用 gin engine，注入 mock service
func setupRouter(svc rolehandler.RoleServiceInterface) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.Locale())
	rh := rolehandler.NewRoleHandler(svc)
	r.GET("/roles", rh.List)
	r.GET("/roles/:id", rh.GetByID)
	r.POST("/roles", rh.Create)
	r.PUT("/roles/:id", rh.Update)
	r.DELETE("/roles/:id", rh.Delete)
	return r
}

func jsonBody(v any) *bytes.Buffer {
	b, _ := json.Marshal(v)
	return bytes.NewBuffer(b)
}

// --- List ---
func TestRoleHandler_List(t *testing.T) {
	cases := []struct {
		name       string
		query      string
		mockRoles  []model.Role
		mockTotal  int64
		mockErr    error
		wantStatus int
		wantTotal  *int64
	}{
		{
			name:       "成功回傳列表",
			query:      "?page=1&perPage=2",
			mockRoles:  []model.Role{{Name: "admin"}, {Name: "editor"}},
			mockTotal:  2,
			wantStatus: http.StatusOK,
			wantTotal:  ptr[int64](2),
		},
		{
			name:       "未帶參數時套用預設值",
			query:      "",
			mockRoles:  []model.Role{},
			mockTotal:  0,
			wantStatus: http.StatusOK,
			wantTotal:  ptr[int64](0),
		},
		{
			name:       "page 為負數回傳 422",
			query:      "?page=-1",
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name:       "service 錯誤回傳 500",
			query:      "?page=1",
			mockErr:    fmt.Errorf("db error"),
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc := &mockRoleService{
				listFn: func(page, perPage int, name string) ([]model.Role, int64, error) {
					return tc.mockRoles, tc.mockTotal, tc.mockErr
				},
			}
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/roles"+tc.query, nil)
			setupRouter(svc).ServeHTTP(w, req)

			assert.Equal(t, tc.wantStatus, w.Code)
			if tc.wantTotal != nil {
				var body map[string]any
				_ = json.Unmarshal(w.Body.Bytes(), &body)
				pagination := body["pagination"].(map[string]any)
				assert.Equal(t, *tc.wantTotal, int64(pagination["total"].(float64)))
			}
		})
	}
}

// --- GetByID ---
func TestRoleHandler_GetByID(t *testing.T) {
	cases := []struct {
		name       string
		pathID     string
		mockRole   model.Role
		mockErr    error
		wantStatus int
	}{
		{
			name:       "成功回傳單筆",
			pathID:     "1",
			mockRole:   model.Role{Name: "admin"},
			wantStatus: http.StatusOK,
		},
		{
			name:       "ID 非數字回傳 404",
			pathID:     "abc",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "找不到回傳 404",
			pathID:     "99",
			mockErr:    gorm.ErrRecordNotFound,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "service 錯誤回傳 500",
			pathID:     "1",
			mockErr:    fmt.Errorf("db error"),
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc := &mockRoleService{
				getByIDFn: func(id uint) (model.Role, error) {
					return tc.mockRole, tc.mockErr
				},
			}
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/roles/"+tc.pathID, nil)
			setupRouter(svc).ServeHTTP(w, req)
			assert.Equal(t, tc.wantStatus, w.Code)
		})
	}
}

// --- Create ---
func TestRoleHandler_Create(t *testing.T) {
	cases := []struct {
		name       string
		body       any
		mockRole   model.Role
		mockErr    error
		wantStatus int
		wantCode   string
	}{
		{
			name:       "成功建立",
			body:       map[string]any{"name": "admin"},
			mockRole:   model.Role{Name: "admin"},
			wantStatus: http.StatusOK,
		},
		{
			name:       "name 為空回傳 422",
			body:       map[string]any{"name": ""},
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name:       "name 超過 20 字回傳 422",
			body:       map[string]any{"name": "123456789012345678901"},
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name:       "名稱重複回傳 422 帶 ROLE_NAME_TAKEN",
			body:       map[string]any{"name": "admin"},
			mockErr:    rolesvc.ErrRoleNameTaken,
			wantStatus: http.StatusUnprocessableEntity,
			wantCode:   "ROLE_NAME_TAKEN",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc := &mockRoleService{
				createFn: func(name string) (model.Role, error) {
					return tc.mockRole, tc.mockErr
				},
			}
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/roles", jsonBody(tc.body))
			req.Header.Set("Content-Type", "application/json")
			setupRouter(svc).ServeHTTP(w, req)

			assert.Equal(t, tc.wantStatus, w.Code)
			if tc.wantCode != "" {
				var body map[string]any
				_ = json.Unmarshal(w.Body.Bytes(), &body)
				assert.Equal(t, tc.wantCode, body["code"])
			}
		})
	}
}

// --- Update ---
func TestRoleHandler_Update(t *testing.T) {
	cases := []struct {
		name       string
		pathID     string
		body       any
		mockRole   model.Role
		mockErr    error
		wantStatus int
		wantCode   string
	}{
		{
			name:       "成功更新",
			pathID:     "1",
			body:       map[string]any{"name": "editor"},
			mockRole:   model.Role{Name: "editor"},
			wantStatus: http.StatusOK,
		},
		{
			name:       "ID 非數字回傳 404",
			pathID:     "abc",
			body:       map[string]any{"name": "editor"},
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "找不到回傳 404",
			pathID:     "99",
			body:       map[string]any{"name": "editor"},
			mockErr:    gorm.ErrRecordNotFound,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "名稱重複回傳 422 帶 ROLE_NAME_TAKEN",
			pathID:     "1",
			body:       map[string]any{"name": "admin"},
			mockErr:    rolesvc.ErrRoleNameTaken,
			wantStatus: http.StatusUnprocessableEntity,
			wantCode:   "ROLE_NAME_TAKEN",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc := &mockRoleService{
				updateFn: func(id uint, name string) (model.Role, error) {
					return tc.mockRole, tc.mockErr
				},
			}
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPut, "/roles/"+tc.pathID, jsonBody(tc.body))
			req.Header.Set("Content-Type", "application/json")
			setupRouter(svc).ServeHTTP(w, req)

			assert.Equal(t, tc.wantStatus, w.Code)
			if tc.wantCode != "" {
				var body map[string]any
				_ = json.Unmarshal(w.Body.Bytes(), &body)
				assert.Equal(t, tc.wantCode, body["code"])
			}
		})
	}
}

// --- Delete ---
func TestRoleHandler_Delete(t *testing.T) {
	cases := []struct {
		name       string
		pathID     string
		mockErr    error
		wantStatus int
		wantCode   string
	}{
		{
			name:       "成功刪除",
			pathID:     "1",
			wantStatus: http.StatusOK,
		},
		{
			name:       "ID 非數字回傳 404",
			pathID:     "abc",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "找不到回傳 404",
			pathID:     "99",
			mockErr:    gorm.ErrRecordNotFound,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "角色使用中回傳 422 帶 ROLE_IN_USE",
			pathID:     "1",
			mockErr:    rolesvc.ErrRoleInUse,
			wantStatus: http.StatusUnprocessableEntity,
			wantCode:   "ROLE_IN_USE",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc := &mockRoleService{
				deleteFn: func(id uint) error {
					return tc.mockErr
				},
			}
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, "/roles/"+tc.pathID, nil)
			setupRouter(svc).ServeHTTP(w, req)

			assert.Equal(t, tc.wantStatus, w.Code)
			if tc.wantCode != "" {
				var body map[string]any
				_ = json.Unmarshal(w.Body.Bytes(), &body)
				assert.Equal(t, tc.wantCode, body["code"])
			}
		})
	}
}

// ptr 回傳任意型別的指標，用於 table-driven test 中區分「零值」和「未設定」
func ptr[T any](v T) *T { return &v }
