package creaturefilters

import (
	"fmt"

	"game-backend/repository"
	"game-backend/repository/filters"
)

// HabitatFactory 棲息地 filter（精確比對）
//
// 使用範例：
//
//	"habitat": "草原"  → WHERE habitat = '草原'
func HabitatFactory(v any) (repository.FilterInterface, error) {
	s, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("habitat filter: value 必須是 string")
	}
	return filters.NewEqFilter("habitat", s), nil
}
