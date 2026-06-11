// Package seeder 負責寫入初始資料（類似 Laravel 的 database/seeders）
// 用途：開發環境填入預設生物、玩家等測試資料
// 執行順序：RunAll → 依序呼叫各資源的 seeder
package seeder

import (
	"game-backend/db"
	"game-backend/logger"
)

// RunAll 依序執行所有 seeder
// 在 main.go 呼叫，於 migration 之後、伺服器啟動之前執行
func RunAll(log logger.Logger) {
	log.Info("開始執行 Seeder...")

	SeedPermissions(db.DB, log)

	log.Info("Seeder 執行完畢")
}
