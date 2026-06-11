package dto

const (
	DefaultPage    = 1
	DefaultPerPage = 20
)

// PaginationRequest 分頁查詢參數（供 List 類 request struct 嵌入使用）
type PaginationRequest struct {
	// 頁碼（最小 1，預設 1）
	Page int `form:"page" binding:"omitempty,positiveInt" example:"1"`
	// 每頁筆數（最小 1，預設 20）
	PerPage int `form:"perPage" binding:"omitempty,positiveInt" example:"20"`
}

// ApplyDefaults 補上未提供時的預設值
func (p *PaginationRequest) ApplyDefaults() {
	if p.Page == 0 {
		p.Page = DefaultPage
	}
	if p.PerPage == 0 {
		p.PerPage = DefaultPerPage
	}
}

// FieldLocaleKeys 回傳欄位名稱 → locale key 的對應，供驗證錯誤翻譯使用
func (PaginationRequest) FieldLocaleKeys() map[string]string {
	return map[string]string{
		"Page":    "page",
		"PerPage": "perPage",
	}
}
