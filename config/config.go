package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config 對應 .env 的所有設定
type Config struct {
	DBHost          string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBName          string
	DBSSLMode       string
	ServerPort      string
	JWTSecret       string
	SwaggerTestHost string // Swagger UI 測試站切換用的主機名稱（不含 scheme）
}

// App 是全域設定實例，載入後直接用 config.App.DBHost 存取
var App *Config

// Load 讀取 .env 並初始化 App
func Load() {
	_ = godotenv.Load() // 找不到 .env 屬預期情境，直接讀系統環境變數

	App = &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "game_db"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		ServerPort:      getEnv("SERVER_PORT", "8000"),
		JWTSecret:       getEnv("JWT_SECRET", ""),
		SwaggerTestHost: getEnv("SWAGGER_TEST_HOST", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
