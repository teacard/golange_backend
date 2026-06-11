package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// ParseID 從 URL path 參數 `:id` 解析為 uint
func ParseID(ctx *gin.Context) (uint, error) {
	n, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(n), nil
}
