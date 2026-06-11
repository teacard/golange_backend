package roleresponse

import "game-backend/dto"

// RoleListResponse 角色列表回應
type RoleListResponse struct {
	// 角色列表
	Data []RoleResponse `json:"data"`
	// 分頁資訊
	Pagination dto.PaginationMeta `json:"pagination"`
}
