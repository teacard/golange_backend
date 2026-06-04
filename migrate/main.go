package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"game-backend/config"
	gamedb "game-backend/db"
	"game-backend/logger"
	"game-backend/seeder"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	config.Load()
	log := logger.New()

	cmd := os.Args[1]

	// db:seed 不需要 migrate 實例
	if cmd == "db:seed" {
		runSeed(log)
		return
	}

	// 其他指令都需要 migrate 實例
	pgURL := buildPGURL()
	m, err := migrate.New("file://migrations", "postgres://"+pgURL)
	if err != nil {
		log.Fatal("Migration 初始化失敗：", err)
	}
	defer func() { _, _ = m.Close() }()

	switch cmd {
	case "migrate:up":
		if err := m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Info("Migration：資料表已是最新版本，無需更新")
				return
			}
			log.Fatal("migrate:up 失敗：", err)
		}
		version, _, vErr := m.Version()
		if vErr != nil {
			log.Info("migrate:up 完成，版本查詢失敗")
		} else {
			log.Infof("migrate:up 完成，目前版本：%d", version)
		}

	case "migrate:rollback":
		if err := m.Steps(-1); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Info("Migration：沒有可回滾的版本")
				return
			}
			log.Fatal("migrate:rollback 失敗：", err)
		}
		version, _, vErr := m.Version()
		switch {
		case vErr == nil:
			log.Infof("migrate:rollback 完成，目前版本：%d", version)
		case errors.Is(vErr, migrate.ErrNilVersion):
			log.Info("migrate:rollback 完成，目前版本：無（已全部回滾）")
		default:
			log.Info("migrate:rollback 完成，版本查詢失敗")
		}

	case "migrate:down":
		if err := m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Info("Migration：已是最舊版本，無需操作")
				return
			}
			log.Fatal("migrate:down 失敗：", err)
		}
		log.Info("migrate:down 完成，所有 migration 已回滾")

	case "migrate:fresh":
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal("migrate:fresh Down 失敗：", err)
		}
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal("migrate:fresh Up 失敗：", err)
		}
		version, _, vErr := m.Version()
		if vErr != nil {
			log.Info("migrate:fresh 完成，版本查詢失敗")
		} else {
			log.Infof("migrate:fresh 完成，目前版本：%d", version)
		}

	case "migrate:version":
		version, dirty, err := m.Version()
		if err != nil {
			if errors.Is(err, migrate.ErrNilVersion) {
				log.Info("目前版本：無（尚未執行任何 migration）")
				return
			}
			log.Fatal("取得版本失敗：", err)
		}
		log.Infof("目前版本：%d，dirty：%v", version, dirty)

	case "migrate:force":
		if len(os.Args) < 3 {
			log.Fatal("用法：go run ./migrate migrate:force <版本號>")
		}
		v, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("版本號格式錯誤：%s", os.Args[2])
		}
		log.Warn("警告：migrate:force 會強制覆寫版本號，請確認 schema 狀態正確再執行")
		if err := m.Force(v); err != nil {
			log.Fatal("migrate:force 失敗：", err)
		}
		log.Infof("migrate:force 完成，強制設定版本為：%d", v)

	default:
		printUsage()
		os.Exit(1)
	}
}

// runSeed 建立 GORM 連線並執行所有 seeder
func runSeed(log logger.Logger) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.App.DBHost,
		config.App.DBPort,
		config.App.DBUser,
		config.App.DBPassword,
		config.App.DBName,
		config.App.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("資料庫連線失敗：", err)
	}
	log.Info("資料庫連線成功")

	gamedb.DB = db
	seeder.RunAll(log)
}

// buildPGURL 組出 migrate 用的 postgres URL（不含 scheme 前綴）
func buildPGURL() string {
	userInfo := url.UserPassword(config.App.DBUser, config.App.DBPassword)
	return fmt.Sprintf(
		"%s@%s:%s/%s?sslmode=%s",
		userInfo.String(),
		config.App.DBHost,
		config.App.DBPort,
		config.App.DBName,
		config.App.DBSSLMode,
	)
}

// printUsage 印出可用指令說明
func printUsage() {
	fmt.Println("用法：go run ./migrate <指令>")
	fmt.Println()
	fmt.Println("可用指令：")
	fmt.Println("  migrate:up          執行所有待執行的 migration")
	fmt.Println("  migrate:rollback    回滾最近一個 migration")
	fmt.Println("  migrate:down        回滾所有 migration")
	fmt.Println("  migrate:fresh       回滾所有再重新執行（清空重建）")
	fmt.Println("  migrate:version     顯示目前 migration 版本")
	fmt.Println("  migrate:force <v>   強制設定版本號（修復 dirty 狀態用）")
	fmt.Println("  db:seed             執行所有 seeder 填入初始資料")
}
