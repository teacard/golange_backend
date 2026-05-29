package repository

import (
	"fmt"
	"reflect"

	"game-backend/repository/filters"
)

// CommonFilters 通用 filter 清單，所有 repository 都可以使用
// key 使用 snake_case，與 SQL 欄位名稱風格一致
var CommonFilters = FilterFactory{

	// id：支援單一值（WHERE id = ?）或 slice（WHERE id IN (...)）
	// 使用範例：
	//   "id": 1           → WHERE id = 1
	//   "id": []int{1,2}  → WHERE id IN (1, 2)
	"id": func(v any) (FilterInterface, error) {
		// reflect.ValueOf 取得 v 的反射資訊
		// reflect.Slice 判斷是否為 slice 型別（[]int、[]uint 等都算）
		if reflect.ValueOf(v).Kind() == reflect.Slice {
			return filters.NewInFilter("id", v), nil
		}
		return filters.NewEqFilter("id", v), nil
	},

	// order_by：排序，值為 "欄位名 ASC/DESC"
	// 使用範例："order_by": "created_at DESC"
	"order_by": func(v any) (FilterInterface, error) {
		s, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("order_by: value 必須是 string，例如 \"created_at DESC\"")
		}
		return filters.NewOrderByFilter(s), nil
	},

	// limit：限制回傳筆數
	// 使用範例："limit": 10
	"limit": func(v any) (FilterInterface, error) {
		n, ok := v.(int)
		if !ok {
			return nil, fmt.Errorf("limit: value 必須是 int")
		}
		return filters.NewLimitFilter(n), nil
	},

	// offset：跳過幾筆（搭配 limit 做分頁）
	// 使用範例："offset": 20  → 跳過前 20 筆，從第 21 筆開始
	"offset": func(v any) (FilterInterface, error) {
		n, ok := v.(int)
		if !ok {
			return nil, fmt.Errorf("offset: value 必須是 int")
		}
		return filters.NewOffsetFilter(n), nil
	},
}

// Resolve 將 filter key + value 解析成 FilterInterface
//
// 解析順序：
//  1. 先找 CommonFilters（通用）
//  2. 找不到再依序找 extraFactories（model 專屬）
//  3. 都找不到 → 回傳 error，明確告知哪個 key 未被註冊
//
// 這樣設計的好處：
//   → service 層只需傳 map，不需要知道 filter 的實作細節
//   → 新增 filter 只需在對應的 FilterFactory 裡加一筆，不改呼叫端
func Resolve(key string, value any, extraFactories ...FilterFactory) (FilterInterface, error) {
	// 1. 通用 filters
	if factory, ok := CommonFilters[key]; ok {
		return factory(value)
	}

	// 2. model 專屬 filters（依序找，先找到就用）
	for _, extra := range extraFactories {
		if factory, ok := extra[key]; ok {
			return factory(value)
		}
	}

	// 3. 未註冊
	return nil, fmt.Errorf(
		"filter key '%s' 未被註冊，請在 CommonFilters 或 model 專屬 FilterFactory 中新增",
		key,
	)
}
