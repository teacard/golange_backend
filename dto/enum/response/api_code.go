package enumresponse

// ApiCodeItem 單筆錯誤碼對照資料
type ApiCodeItem struct {
	// 錯誤碼值（機器可讀）
	Value string `json:"value" example:"ROLE_NAME_TAKEN"`
	// 對應中文說明（人類可讀）
	Label string `json:"label" example:"角色名稱已存在"`
}
