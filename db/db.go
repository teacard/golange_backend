package db

import (
	"errors"
	"fmt"
	"game-backend/config"
	"game-backend/logger"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // file source

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB 是全域 GORM 連線實例
var DB *gorm.DB

// Init 連接 PostgreSQL 並執行版本化 migration
func Init(log logger.Logger) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.App.DBHost,
		config.App.DBPort,
		config.App.DBUser,
		config.App.DBPassword,
		config.App.DBName,
		config.App.DBSSLMode,
	)

	// ── 1. 建立 GORM 連線 ──────────────────────────────────────
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("資料庫連線失敗：", err)
	}
	log.Info("資料庫連線成功")

	// ── 2. 執行 golang-migrate（SQL 版本化 migration）──────────
	runMigrations(log)
}

// runMigrations 讀取 migrations/ 目錄內的 SQL 檔案並執行到最新版本
func runMigrations(log logger.Logger) {
	// file source 路徑（相對於執行目錄）
	m, err := migrate.New("file://migrations", "postgres://"+buildPGURL())
	if err != nil {
		log.Fatal("Migration 初始化失敗：", err)
	}
	defer func() { _, _ = m.Close() }()

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info("Migration：資料表已是最新版本，無需更新")
			return
		}
		log.Fatal("Migration 執行失敗：", err)
	}

	version, _, _ := m.Version()
	log.Infof("Migration 完成，目前版本：%d", version)
}

// buildPGURL 組出 migrate 用的 postgres URL（不含 scheme 前綴）
func buildPGURL() string {
	return fmt.Sprintf(
		"%s:%s@%s:%s/%s?sslmode=%s",
		config.App.DBUser,
		config.App.DBPassword,
		config.App.DBHost,
		config.App.DBPort,
		config.App.DBName,
		config.App.DBSSLMode,
	)
}
