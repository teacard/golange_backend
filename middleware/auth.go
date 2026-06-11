package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"game-backend/config"
	"game-backend/exception"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Auth 驗證 Bearer JWT，無效或缺少 token 直接回 401 並中止請求
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exception.BaseErrorResponse{
				Error: "未提供授權 Token",
			})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			// 確認簽名演算法是 HMAC，避免 alg:none 攻擊
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("非預期的簽名演算法：%v", t.Header["alg"])
			}
			return []byte(config.App.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exception.BaseErrorResponse{
				Error: "Token 無效或已過期",
			})
			return
		}

		ctx.Next()
	}
}
