package model

import (
	"time"

	"gorm.io/gorm"
)

// Inventory 對應資料庫的 inventories 資料表
// 記錄每位玩家已蒐集的生物（背包）
//
// FK 設計：
//   → PlayerID 儲存的是 players.id（uint，DB primary key）
//   → CreatureID 儲存的是 creatures.id（uint，DB primary key）
//   → 全系統統一用 DB primary key 做關聯，不混用 business id（字串）
//
// 注意：API 層查詢背包時，應從認證 Token 取得玩家身份
//   → 不應該讓前端直接傳 player_id 查詢，避免越權存取他人背包
//   → 認證 middleware 之後會處理這件事
type Inventory struct {
	gorm.Model

	// PlayerID 對應 players.id（DB primary key，uint）
	// 對應 SQL：player_id INTEGER NOT NULL REFERENCES players(id)
	PlayerID uint

	// CreatureID 對應 creatures.id（DB primary key，uint）
	// 對應 SQL：creature_id INTEGER NOT NULL REFERENCES creatures(id)
	CreatureID uint

	// CaughtAt 捕獲時間
	// time.Time 對應 SQL 的 TIMESTAMPTZ
	CaughtAt time.Time
}
