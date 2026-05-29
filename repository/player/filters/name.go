package playerfilters

import (
	"fmt"

	"game-backend/repository"
	"game-backend/repository/filters"
)

// NameFactory 玩家名稱 filter（精確比對）
//
// 使用範例：
//
//	"name": "chen"  → WHERE name = 'chen'
func NameFactory(v any) (repository.FilterInterface, error) {
	s, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("name filter: value 必須是 string")
	}
	return filters.NewEqFilter("name", s), nil
}
