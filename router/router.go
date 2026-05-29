package router

import (
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	// TODO: 在這裡登記路由
	_ = r.Group("/api")

	return r
}
