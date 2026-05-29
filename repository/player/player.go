package player

import (
	"game-backend/model"
	"game-backend/repository"
	playerfilters "game-backend/repository/player/filters"

	"gorm.io/gorm"
)

// PlayerRepository 玩家專屬 Repository
// 內嵌 BaseRepository，並注入玩家專屬 FilterFactory
type PlayerRepository struct {
	repository.BaseRepository[model.Player]
}

// playerFilters 玩家專屬 filter 清單
var playerFilters = repository.FilterFactory{
	"name": playerfilters.NameFactory,
}

func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{
		BaseRepository: repository.BaseRepository[model.Player]{DB: db},
	}
}

// FindAll 覆寫 BaseRepository.FindAll，自動帶入 playerFilters
//
// 使用範例（在 service 層）：
//
//	players, err := repo.FindAll(map[string]any{
//	    "name":     "chen",
//	    "order_by": "created_at DESC",
//	})
func (r *PlayerRepository) FindAll(conditions map[string]any) ([]model.Player, error) {
	return r.BaseRepository.FindAll(conditions, playerFilters)
}

// FindOne 覆寫 BaseRepository.FindOne，自動帶入 playerFilters
func (r *PlayerRepository) FindOne(conditions map[string]any) (model.Player, error) {
	return r.BaseRepository.FindOne(conditions, playerFilters)
}
