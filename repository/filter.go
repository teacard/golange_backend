package repository

import "gorm.io/gorm"

// FilterInterface 每個 filter 只需實作 Apply
// Apply 接收目前的 query，加上自己的條件後回傳新的 query
// GORM 的 *gorm.DB 是 chain 設計：每次 Where/Order/Limit 都回傳新的 DB 物件，不修改原本的
type FilterInterface interface {
	Apply(db *gorm.DB) *gorm.DB
}

// FilterFactory 是一個 map，key 是 filter 名稱（snake_case），value 是建立 FilterInterface 的函式
//
// 解析順序（在 Resolve 函式中執行）：
//  1. 先找 CommonFilters（通用，所有 repository 共用）
//  2. 找不到再找 model 專屬的 FilterFactory
//  3. 都找不到 → 回傳 error
//
// 使用範例（在 model 專屬 repository 中定義）：
//
//	var CreatureFilters = FilterFactory{
//	    "rarity":  creaturefilters.RarityFactory,
//	    "habitat": creaturefilters.HabitatFactory,
//	}
type FilterFactory map[string]func(value any) (FilterInterface, error)
