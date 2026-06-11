package role

import (
	"errors"
	"net/http"

	"game-backend/dto"
	rolerequest "game-backend/dto/role/request"
	roleresponse "game-backend/dto/role/response"
	api_code "game-backend/enum/api_code"
	"game-backend/handler"
	"game-backend/model"
	rolesvc "game-backend/service/role"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RoleServiceInterface 讓測試可以注入 mock
type RoleServiceInterface interface {
	List(page, perPage int, name string) ([]model.Role, int64, error)
	GetByID(id uint) (model.Role, error)
	Create(name string) (model.Role, error)
	Update(id uint, name string) (model.Role, error)
	Delete(id uint) error
}

// RoleHandler 角色 Handler
type RoleHandler struct {
	svc RoleServiceInterface
}

// NewRoleHandler 建立 RoleHandler，注入 service 依賴
func NewRoleHandler(svc RoleServiceInterface) *RoleHandler {
	return &RoleHandler{svc: svc}
}

// List 取得角色列表，支援名稱模糊搜尋與分頁
func (rh *RoleHandler) List(ctx *gin.Context) {
	var req rolerequest.ListRoleRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		handler.RespondValidationError(ctx, err, req.FieldLocaleKeys())
		return
	}
	req.ApplyDefaults()

	roles, total, err := rh.svc.List(req.Page, req.PerPage, req.Name)
	if err != nil {
		handler.RespondInternalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, roleresponse.RoleListResponse{
		Data: dto.MapSlice(roles, roleresponse.FromModel),
		Pagination: dto.PaginationMeta{
			Page:    req.Page,
			PerPage: req.PerPage,
			Total:   total,
		},
	})
}

// GetByID 取得單筆角色
func (rh *RoleHandler) GetByID(ctx *gin.Context) {
	id, err := handler.ParseID(ctx)
	if err != nil {
		handler.RespondNotFound(ctx, "role.not_found")
		return
	}

	role, err := rh.svc.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handler.RespondNotFound(ctx, "role.not_found")
			return
		}
		handler.RespondInternalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, roleresponse.FromModel(role))
}

// Create 建立新角色
func (rh *RoleHandler) Create(ctx *gin.Context) {
	var req rolerequest.CreateRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handler.RespondValidationError(ctx, err, req.FieldLocaleKeys())
		return
	}

	role, err := rh.svc.Create(req.Name)
	if err != nil {
		if errors.Is(err, rolesvc.ErrRoleNameTaken) {
			handler.RespondUnprocessableError(ctx, api_code.RoleNameTaken)
			return
		}
		handler.RespondInternalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, roleresponse.FromModel(role))
}

// Update 更新角色名稱
func (rh *RoleHandler) Update(ctx *gin.Context) {
	id, err := handler.ParseID(ctx)
	if err != nil {
		handler.RespondNotFound(ctx, "role.not_found")
		return
	}

	var req rolerequest.UpdateRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handler.RespondValidationError(ctx, err, req.FieldLocaleKeys())
		return
	}

	role, err := rh.svc.Update(id, req.Name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handler.RespondNotFound(ctx, "role.not_found")
			return
		}
		if errors.Is(err, rolesvc.ErrRoleNameTaken) {
			handler.RespondUnprocessableError(ctx, api_code.RoleNameTaken)
			return
		}
		handler.RespondInternalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, roleresponse.FromModel(role))
}

// Delete 刪除角色
func (rh *RoleHandler) Delete(ctx *gin.Context) {
	id, err := handler.ParseID(ctx)
	if err != nil {
		handler.RespondNotFound(ctx, "role.not_found")
		return
	}

	if err := rh.svc.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handler.RespondNotFound(ctx, "role.not_found")
			return
		}
		if errors.Is(err, rolesvc.ErrRoleInUse) {
			handler.RespondUnprocessableError(ctx, api_code.RoleInUse)
			return
		}
		handler.RespondInternalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
