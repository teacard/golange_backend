package role

import (
	"errors"
	"strings"

	"game-backend/model"
	rolerepo "game-backend/repository/role"

	"gorm.io/gorm"
)

var (
	ErrRoleNameTaken = errors.New("role name taken") // 角色名稱重複
	ErrRoleInUse     = errors.New("role in use")     // 角色仍有帳號使用中，不能刪除
)

// RoleService 角色業務邏輯
type RoleService struct {
	repo *rolerepo.RoleRepository
	db   *gorm.DB
}

// NewRoleService 建立 RoleService，注入 db 依賴
func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{
		repo: rolerepo.NewRoleRepository(db),
		db:   db,
	}
}

// List 取得角色列表，支援名稱模糊搜尋與分頁
func (rs *RoleService) List(page, perPage int, name string) ([]model.Role, int64, error) {
	conditions := map[string]any{}
	if name != "" {
		conditions["name"] = name
	}

	result, err := rs.repo.Paginate(page, perPage, conditions)
	if err != nil {
		return nil, 0, err
	}

	return result.Data, result.Total, nil
}

// GetByID 取得單筆角色，找不到回傳 gorm.ErrRecordNotFound
func (rs *RoleService) GetByID(id uint) (model.Role, error) {
	return rs.repo.FindOne(map[string]any{"id": id})
}

// Create 建立新角色，名稱重複回傳 ErrRoleNameTaken
func (rs *RoleService) Create(name string) (model.Role, error) {
	var nameCount int64
	if err := rs.db.Model(&model.Role{}).Where("name = ?", name).Count(&nameCount).Error; err != nil {
		return model.Role{}, err
	}
	if nameCount > 0 {
		return model.Role{}, ErrRoleNameTaken
	}

	role := model.Role{Name: name}
	if err := rs.repo.Create(&role); err != nil {
		if isDuplicateKeyError(err) {
			return model.Role{}, ErrRoleNameTaken
		}
		return model.Role{}, err
	}
	return role, nil
}

// Update 更新角色名稱，名稱未變動略過唯一性查詢
func (rs *RoleService) Update(id uint, name string) (model.Role, error) {
	role, err := rs.repo.FindOne(map[string]any{"id": id})
	if err != nil {
		return model.Role{}, err
	}

	if role.Name != name {
		var nameCount int64
		if err := rs.db.Model(&model.Role{}).Where("name = ?", name).Count(&nameCount).Error; err != nil {
			return model.Role{}, err
		}
		if nameCount > 0 {
			return model.Role{}, ErrRoleNameTaken
		}
	}

	role.Name = name
	if err := rs.repo.Save(&role); err != nil {
		if isDuplicateKeyError(err) {
			return model.Role{}, ErrRoleNameTaken
		}
		return model.Role{}, err
	}
	return role, nil
}

// Delete 刪除角色，有帳號使用時回傳 ErrRoleInUse
func (rs *RoleService) Delete(id uint) error {
	role, err := rs.repo.FindOne(map[string]any{"id": id})
	if err != nil {
		return err
	}

	var accountCount int64
	if err := rs.db.Model(&model.ModelRole{}).Where("role_id = ?", id).Count(&accountCount).Error; err != nil {
		return err
	}
	if accountCount > 0 {
		return ErrRoleInUse
	}

	return rs.repo.Delete(&role)
}

// isDuplicateKeyError 判斷是否為 PostgreSQL 唯一鍵衝突錯誤（23505 是 PG 的 unique_violation 錯誤碼）
func isDuplicateKeyError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "23505")
}
