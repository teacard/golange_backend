package creaturefilters

import (
	"fmt"
	"reflect"

	"game-backend/repository"
	"game-backend/repository/filters"
)

// RarityFactory 稀有度 filter
// 支援單一值（WHERE rarity = ?）或 slice（WHERE rarity IN (...)）
//
// 使用範例：
//
//	"rarity": "uncommon"                       → WHERE rarity = 'uncommon'
//	"rarity": []string{"common", "uncommon"}   → WHERE rarity IN ('common', 'uncommon')
func RarityFactory(v any) (repository.FilterInterface, error) {
	if reflect.ValueOf(v).Kind() == reflect.Slice {
		return filters.NewInFilter("rarity", v), nil
	}
	s, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("rarity filter: value 必須是 string 或 []string")
	}
	return filters.NewEqFilter("rarity", s), nil
}
