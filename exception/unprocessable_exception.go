package exception

import api_code "game-backend/enum/api_code"

// UnprocessableEntityResponse 422 業務驗證失敗，帶機器可讀錯誤碼
type UnprocessableEntityResponse struct {
	BaseErrorResponse
	// 業務錯誤碼，前端可根據此碼做對應處理
	Code api_code.ApiCode `json:"code" example:"ROLE_NAME_TAKEN"`
}
