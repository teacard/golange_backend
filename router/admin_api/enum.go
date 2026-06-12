package adminapi

import (
	enumhandler "game-backend/handler/enum"

	"github.com/gin-gonic/gin"
)

func init() {
	RegisterPublic("/enums", registerEnumRoutes)
}

func registerEnumRoutes(rg *gin.RouterGroup) {
	eh := enumhandler.NewEnumHandler()
	// 錯誤碼對照表：前端用來將 ApiCode 轉成中文說明，不需登入即可取得
	rg.GET("/api-code", eh.ListApiCodes)
}
