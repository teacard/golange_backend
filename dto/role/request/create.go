package rolerequest

// CreateRoleRequest 建立角色請求
type CreateRoleRequest struct {
	// 角色名稱（最多 20 字）
	Name string `json:"name" binding:"required,max=20" example:"管理員"`
}

// FieldLocaleKeys 回傳欄位名稱 → locale key 的對應，供驗證錯誤翻譯使用
func (CreateRoleRequest) FieldLocaleKeys() map[string]string {
	return map[string]string{
		"Name": "name",
	}
}
