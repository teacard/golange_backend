package middleware

import (
	"game-backend/locale"

	"github.com/gin-gonic/gin"
)

// LocalizerKey 用 const 避免字串打錯造成 ctx.Get 拿不到值的 silent bug
const LocalizerKey = "localizer"

// Locale 從 Accept-Language header 偵測語系，預設 zh-TW
func Locale() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		acceptLanguage := ctx.GetHeader("Accept-Language")
		ctx.Set(LocalizerKey, locale.Localizer(acceptLanguage))
		ctx.Next()
	}
}
