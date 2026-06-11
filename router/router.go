package router

import (
	"net/http"

	"game-backend/db"
	enumhandler "game-backend/handler/enum"
	rolehandler "game-backend/handler/role"
	"game-backend/middleware"
	rolesvc "game-backend/service/role"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Locale())

	// Swagger UI：http://localhost:8000/swagger/admin-api（後台）
	// 未來前台：http://localhost:8000/swagger/api
	adminSwaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	r.GET("/swagger/admin-api/*any", func(c *gin.Context) {
		if c.Param("any") == "/" {
			c.Redirect(http.StatusMovedPermanently, "/swagger/admin-api/index.html")
			return
		}
		adminSwaggerHandler(c)
	})

	adminAPI := r.Group("/admin-api", middleware.Auth())
	{
		roles := adminAPI.Group("/roles")
		rh := rolehandler.NewRoleHandler(rolesvc.NewRoleService(db.DB))
		roles.GET("", rh.List)
		roles.GET("/:id", rh.GetByID)
		roles.POST("", rh.Create)
		roles.PUT("/:id", rh.Update)
		roles.DELETE("/:id", rh.Delete)

		enums := adminAPI.Group("/enums")
		eh := enumhandler.NewEnumHandler()
		enums.GET("/api-code", eh.ListApiCodes)
	}

	return r
}
