package rolerequest

// UpdateRoleRequest 更新角色請求（embed CreateRoleRequest，欄位相同時共用驗證規則）
type UpdateRoleRequest struct {
	CreateRoleRequest
}
