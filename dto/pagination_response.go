package dto

// PaginationMeta 分頁元資訊（回應用）
type PaginationMeta struct {
	// 目前頁碼
	Page int `json:"page" example:"1"`
	// 每頁筆數
	PerPage int `json:"perPage" example:"20"`
	// 符合條件的總筆數
	Total int64 `json:"total" example:"100"`
}
