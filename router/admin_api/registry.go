// Package adminapi 提供 admin-api 路由群組的完整設定與內部自動註冊機制。
//
// 對外：Setup(r) 供上層 router 呼叫，一次掛載所有 /admin-api 路由。
// 對內：各資源檔在 init() 裡呼叫對應的 Register 函式自行登記，
//
//	新增資源不需修改任何現有檔案。
//
// 公開路由用 RegisterPublic()，需登入的路由用 Register()。
package adminapi

import (
	"game-backend/middleware"

	"github.com/gin-gonic/gin"
)

// RouteRegistrar 是路由註冊函式的共同型別
type RouteRegistrar func(rg *gin.RouterGroup)

type routeEntry struct {
	prefix    string
	registrar RouteRegistrar
}

// ── 公開路由 registry ─────────────────────────────────────────────────────────

var publicRegistry []routeEntry

// RegisterPublic 供不需登入的資源檔在 init() 裡呼叫。
func RegisterPublic(prefix string, fn RouteRegistrar) {
	publicRegistry = append(publicRegistry, routeEntry{prefix: prefix, registrar: fn})
}

func registerAllPublic(rg *gin.RouterGroup) {
	for _, entry := range publicRegistry {
		entry.registrar(rg.Group(entry.prefix))
	}
}

// ── 受保護路由 registry ───────────────────────────────────────────────────────

var protectedRegistry []routeEntry

// Register 供需要登入的資源檔在 init() 裡呼叫。
func Register(prefix string, fn RouteRegistrar) {
	protectedRegistry = append(protectedRegistry, routeEntry{prefix: prefix, registrar: fn})
}

func registerAll(rg *gin.RouterGroup) {
	for _, entry := range protectedRegistry {
		entry.registrar(rg.Group(entry.prefix))
	}
}

// ── 對外入口 ──────────────────────────────────────────────────────────────────

// Setup 掛載所有 /admin-api 路由，由上層 router 呼叫。
func Setup(r *gin.Engine) {
	// 公開路由（不需要 JWT Token）
	registerAllPublic(r.Group("/admin-api"))

	// 受保護路由（middleware.Auth() 驗證失敗直接回傳 401）
	registerAll(r.Group("/admin-api", middleware.Auth()))
}
