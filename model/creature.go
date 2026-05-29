package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Creature 對應資料庫的 creatures 資料表
//
// gorm.Model 自動提供四個欄位：
//   - ID        uint             → 主鍵
//   - CreatedAt time.Time        → 建立時間（INSERT 時自動填入）
//   - UpdatedAt time.Time        → 更新時間（UPDATE 時自動填入）
//   - DeletedAt gorm.DeletedAt   → 軟刪除時間（Delete() 時填入，不真正刪除資料列）
//
// 欄位名稱對應規則：Go 的 CamelCase 自動轉成 SQL 的 snake_case
//   例如：CreatureID → creature_id、SpecialParts → special_parts
//
// json tag 不在這裡定義，API 回傳格式由 DTO 負責
type Creature struct {
	gorm.Model

	CreatureID   string
	Rarity       string
	Personality  string
	Habitat      string
	Colors       datatypes.JSON
	SpecialParts datatypes.JSON
	Animation    datatypes.JSON
}
