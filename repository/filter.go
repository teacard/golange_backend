package repository

import (
	"fmt"
	"reflect"

	"game-backend/repository/filters"
	"gorm.io/gorm"
)

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

// CommonFilters 通用 filter 清單，所有 repository 都可以使用
// key 使用 snake_case，與 SQL 欄位名稱風格一致
var CommonFilters = FilterFactory{

	// id：支援單一值（WHERE id = ?）或 slice（WHERE id IN (...)）
	// 使用範例：
	//   "id": 1           → WHERE id = 1
	//   "id": []int{1,2}  → WHERE id IN (1, 2)
	"id": func(v any) (FilterInterface, error) {
		if reflect.ValueOf(v).Kind() == reflect.Slice {
			return filters.NewInFilter("id", v), nil
		}
		return filters.NewEqFilter("id", v), nil
	},

	// order_by：排序，值為 "欄位名 ASC/DESC"
	// 使用範例："order_by": "created_at DESC"
	// value 由 service 層傳入，呼叫端保證型別為 string
	"order_by": func(v any) (FilterInterface, error) {
		return filters.NewOrderByFilter(v.(string)), nil
	},

	// limit：限制回傳筆數
	// 使用範例："limit": 10
	// value 由 service 層傳入，呼叫端保證型別為 int
	"limit": func(v any) (FilterInterface, error) {
		return filters.NewLimitFilter(v.(int)), nil
	},

	// offset：跳過幾筆（搭配 limit 做分頁）
	// 使用範例："offset": 20  → 跳過前 20 筆，從第 21 筆開始
	// value 由 service 層傳入，呼叫端保證型別為 int
	"offset": func(v any) (FilterInterface, error) {
		return filters.NewOffsetFilter(v.(int)), nil
	},
}

// Resolve 將 filter key + value 解析成 FilterInterface
//
// 解析順序：
//  1. 先找 CommonFilters（通用）
//  2. 找不到再依序找 extraFactories（model 專屬）
//  3. 都找不到 → 回傳 error，明確告知哪個 key 未被註冊
func Resolve(key string, value any, extraFactories ...FilterFactory) (FilterInterface, error) {
	if factory, ok := CommonFilters[key]; ok {
		return factory(value)
	}

	for _, extra := range extraFactories {
		if factory, ok := extra[key]; ok {
			return factory(value)
		}
	}

	return nil, fmt.Errorf(
		"filter key '%s' 未被註冊，請在 CommonFilters 或 model 專屬 FilterFactory 中新增",
		key,
	)
}
