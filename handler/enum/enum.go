package enumhandler

import (
	"net/http"

	enumresponse "game-backend/dto/enum/response"
	api_code "game-backend/enum/api_code"
	"game-backend/handler"

	"github.com/gin-gonic/gin"
)

// EnumHandler 提供 enum 對照表查詢
type EnumHandler struct{}

func NewEnumHandler() *EnumHandler {
	return &EnumHandler{}
}

// ListApiCodes 回傳所有業務錯誤碼與對應的本地化說明
func (eh *EnumHandler) ListApiCodes(ctx *gin.Context) {
	items := make([]enumresponse.ApiCodeItem, len(api_code.AllApiCodes))
	for i, code := range api_code.AllApiCodes {
		items[i] = enumresponse.ApiCodeItem{
			Value: string(code),
			Label: handler.Localize(ctx, code.LocaleKey()),
		}
	}
	ctx.JSON(http.StatusOK, items)
}
