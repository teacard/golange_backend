package adminapi

import (
	"game-backend/db"
	rolehandler "game-backend/handler/role"
	rolesvc "game-backend/service/role"

	"github.com/gin-gonic/gin"
)

// init 在 package 被 import 時自動執行，將角色路由登記到 registry。
// 之後新增其他資源（item.go、user.go 等）只要照相同模式寫 init() 即可，
// 不需要修改 admin_api.go 或其他任何檔案。
func init() {
	Register("/roles", registerRoleRoutes)
}

func registerRoleRoutes(rg *gin.RouterGroup) {
	rh := rolehandler.NewRoleHandler(rolesvc.NewRoleService(db.DB))
	rg.GET("", rh.List)
	rg.GET("/:id", rh.GetByID)
	rg.POST("", rh.Create)
	rg.PUT("/:id", rh.Update)
	rg.DELETE("/:id", rh.Delete)
}
