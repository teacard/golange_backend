package role

import (
	"game-backend/model"
	"game-backend/repository"
	rolefilters "game-backend/repository/role/filters"

	"gorm.io/gorm"
)

// RoleRepository 角色專屬 Repository，嵌入 BaseRepository 繼承通用 CRUD
type RoleRepository struct {
	repository.BaseRepository[model.Role]
}

// roleFilters 角色專屬的 filter 清單，傳給 BaseRepository 的各個查詢方法
var roleFilters = repository.FilterFactory{
	"name": rolefilters.NameFactory,
}

// NewRoleRepository 建立 RoleRepository，注入 db 依賴
func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		BaseRepository: repository.BaseRepository[model.Role]{DB: db},
	}
}

// Paginate 分頁查詢，自動帶入角色專屬 filter
func (r *RoleRepository) Paginate(page, perPage int, conditions map[string]any) (repository.PaginationResult[model.Role], error) {
	return r.BaseRepository.Paginate(page, perPage, conditions, roleFilters)
}

// FindAll 查詢多筆角色，自動帶入角色專屬 filter
func (r *RoleRepository) FindAll(conditions map[string]any) ([]model.Role, error) {
	return r.BaseRepository.FindAll(conditions, roleFilters)
}

// FindOne 查詢單筆角色，自動帶入角色專屬 filter
func (r *RoleRepository) FindOne(conditions map[string]any) (model.Role, error) {
	return r.BaseRepository.FindOne(conditions, roleFilters)
}

// Count 計算符合條件的角色筆數，自動帶入角色專屬 filter
func (r *RoleRepository) Count(conditions map[string]any) (int64, error) {
	return r.BaseRepository.Count(conditions, roleFilters)
}

// Exists 確認是否存在符合條件的角色，自動帶入角色專屬 filter
func (r *RoleRepository) Exists(conditions map[string]any) (bool, error) {
	return r.BaseRepository.Exists(conditions, roleFilters)
}
