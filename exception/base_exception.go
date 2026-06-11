package exception

// BaseErrorResponse 通用錯誤回應基底
type BaseErrorResponse struct {
	// 錯誤訊息
	Error string `json:"error" example:"找不到角色"`
}
