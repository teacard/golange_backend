package model

// CreatureLocale 對應資料庫的 creature_locales 資料表
// 儲存生物的多語系文字（名稱、圖鑑說明）
//
// 使用 ChildModel 而非 gorm.Model 的原因：
//   → 這是子表，不需要獨立的 deleted_at
//   → 父表（Creature）軟刪除後，查詢時 JOIN creatures WHERE deleted_at IS NULL 自動過濾
//   → 詳見 base.go 的說明
//
// FK 設計：
//   → CreatureID 儲存的是 creatures.id（uint，DB primary key）
//   → 全系統統一用 DB primary key 做關聯，不混用 business id（字串）
type CreatureLocale struct {
	ChildModel

	// CreatureID 是 creatures 資料表的 DB 主鍵（uint）
	// 對應 SQL：creature_id INTEGER NOT NULL REFERENCES creatures(id)
	CreatureID uint

	// LangCode 語言代碼，遵循 BCP 47 標準
	// 常用值：zh-TW（繁體中文）、en（英文）、ja（日文）
	LangCode string

	Name        string
	Description string
}
