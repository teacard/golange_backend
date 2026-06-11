package roleresponse

import "game-backend/model"

// RoleResponse 單筆角色回應
type RoleResponse struct {
	// 角色 ID
	ID uint `json:"id" example:"1"`
	// 角色名稱
	Name string `json:"name" example:"管理員"`
}

// FromModel 將 model.Role 轉換為 RoleResponse
func FromModel(role model.Role) RoleResponse {
	return RoleResponse{ID: role.ID, Name: role.Name}
}
