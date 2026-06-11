package handler

import (
	"net/http"

	api_code "game-backend/enum/api_code"
	"game-backend/exception"

	"github.com/gin-gonic/gin"
)

// RespondValidationError 回傳 422 並帶入本地化的驗證錯誤訊息
func RespondValidationError(ctx *gin.Context, err error, fieldKeys map[string]string) {
	ctx.JSON(http.StatusUnprocessableEntity, exception.BaseErrorResponse{
		Error: TranslateValidationError(ctx, err, fieldKeys),
	})
}

// RespondNotFound 回傳 404 並帶入本地化訊息
func RespondNotFound(ctx *gin.Context, messageID string) {
	ctx.JSON(http.StatusNotFound, exception.BaseErrorResponse{
		Error: Localize(ctx, messageID),
	})
}

// RespondUnprocessableError 回傳 422 並帶入 ApiCode 與本地化訊息
func RespondUnprocessableError(ctx *gin.Context, code api_code.ApiCode) {
	ctx.JSON(http.StatusUnprocessableEntity, exception.UnprocessableEntityResponse{
		BaseErrorResponse: exception.BaseErrorResponse{Error: Localize(ctx, code.LocaleKey())},
		Code:              code,
	})
}

// RespondInternalError 回傳 500 並帶入原始錯誤訊息
func RespondInternalError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, exception.BaseErrorResponse{
		Error: err.Error(),
	})
}
