package main

import (
	"game-backend/config"
	"game-backend/db"
	_ "game-backend/docs"
	"game-backend/logger"
	"game-backend/router"
	"game-backend/seeder"
)

// @title           Familiar API
// @version         1.0
// @description     Familiar後端 API 文件
// @host            localhost:8000
// @BasePath        /api
func main() {
	// 1. 載入 .env 設定
	config.Load()

	// 2. 初始化 logger
	log := logger.New()

	// 3. 連接資料庫 + 執行 Migration（golang-migrate）
	db.Init(log)

	// 4. 執行 Seeder（寫入初始資料）
	seeder.RunAll(log)

	// 5. 啟動 HTTP 路由
	r := router.Setup()

	log.Infof("伺服器啟動於 http://localhost:%s", config.App.ServerPort)
	if err := r.Run(":" + config.App.ServerPort); err != nil {
		log.Fatal("伺服器啟動失敗：", err)
	}
}
