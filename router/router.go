// Package router 負責建立 HTTP 路由引擎並將所有路由群組組裝在一起。
//
// 各路由群組（swagger、admin-api 等）定義在獨立檔案中，
// 並透過各自的 init() 向 registry 報名。
// Setup() 只負責建立引擎與呼叫 setupAll，不需隨路由增減而修改。
package router

import (
	adminapi "game-backend/router/admin_api"
	"game-backend/middleware"

	"github.com/gin-gonic/gin"
)

// routeSetup 是各路由群組的設定函式型別
type routeSetup func(r *gin.Engine)

// groups 儲存所有透過 registerGroup() 登記的路由群組
var groups []routeSetup

// registerGroup 供各群組檔案在 init() 裡呼叫，將自己登記到 groups。
func registerGroup(fn routeSetup) {
	groups = append(groups, fn)
}

// setupAll 用 for 迴圈依序呼叫所有已登記的路由群組設定函式。
func setupAll(r *gin.Engine) {
	for _, fn := range groups {
		fn(r)
	}
}

func init() {
	// 將各路由群組登記至 registry，往後新增群組在此加一行即可
	registerGroup(adminapi.Setup)
}

// Setup 建立並回傳設定好的 Gin 引擎。
func Setup() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Locale())

	// 所有路由群組已透過 init() 登記至 registry，
	// setupAll 用 for 迴圈統一掛載，不需在此逐一列舉
	setupAll(r)

	return r
}
