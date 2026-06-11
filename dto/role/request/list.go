package rolerequest

import "game-backend/dto"

// ListRoleRequest 角色列表查詢參數
type ListRoleRequest struct {
	dto.PaginationRequest
	// 角色名稱關鍵字（模糊搜尋）
	Name string `form:"name" example:"管理"`
}
