package main

import (
	"game-backend/appvalidator"
	"game-backend/config"
	"game-backend/db"
	_ "game-backend/docs"
	"game-backend/locale"
	"game-backend/logger"
	"game-backend/router"
	"game-backend/seeder"
)

// @title           Familiar API
// @version         1.0
// @description     Familiar後端 API 文件
// @host            localhost:8000
// @BasePath        /
//
// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                 輸入格式：Bearer {token}
func main() {
	// 1. 載入 .env 設定
	config.Load()

	// 2. 初始化 logger
	log := logger.New()

	// 3. 初始化 i18n 語系包
	locale.Init()

	// 4. 註冊自訂 validator tag
	appvalidator.Init()

	// 5. 連接資料庫 + 執行 Migration（golang-migrate）
	db.Init(log)

	// 6. 執行 Seeder（寫入初始資料）
	seeder.RunAll(log)

	// 7. 啟動 HTTP 路由
	r := router.Setup()

	log.Infof("伺服器啟動於 http://localhost:%s", config.App.ServerPort)
	if err := r.Run(":" + config.App.ServerPort); err != nil {
		log.Fatal("伺服器啟動失敗：", err)
	}
}
