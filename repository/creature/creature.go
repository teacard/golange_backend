package creature

import (
	"game-backend/model"
	"game-backend/repository"
	creaturefilters "game-backend/repository/creature/filters"

	"gorm.io/gorm"
)

// CreatureRepository 生物專屬 Repository
// 內嵌 BaseRepository，並注入生物專屬 FilterFactory
//
// 為什麼要覆寫 FindAll / FindOne？
//   → BaseRepository 的 FindAll 需要外部傳入 extraFactories
//   → 覆寫後，service 層只需傳 map，不需要知道 FilterFactory 的存在
//   → 解析順序：CommonFilters → creatureFilters → error
type CreatureRepository struct {
	repository.BaseRepository[model.Creature]
}

// creatureFilters 生物專屬 filter 清單
// 只有在 CommonFilters 找不到 key 時才會進到這裡
var creatureFilters = repository.FilterFactory{
	"rarity":  creaturefilters.RarityFactory,
	"habitat": creaturefilters.HabitatFactory,
}

func NewCreatureRepository(db *gorm.DB) *CreatureRepository {
	return &CreatureRepository{
		BaseRepository: repository.BaseRepository[model.Creature]{DB: db},
	}
}

// FindAll 覆寫 BaseRepository.FindAll，自動帶入 creatureFilters
//
// 使用範例（在 service 層）：
//
//	creatures, err := repo.FindAll(map[string]any{
//	    "rarity":   "uncommon",
//	    "habitat":  "草原",
//	    "order_by": "created_at DESC",
//	    "limit":    10,
//	})
func (r *CreatureRepository) FindAll(conditions map[string]any) ([]model.Creature, error) {
	return r.BaseRepository.FindAll(conditions, creatureFilters)
}

// FindOne 覆寫 BaseRepository.FindOne，自動帶入 creatureFilters
func (r *CreatureRepository) FindOne(conditions map[string]any) (model.Creature, error) {
	return r.BaseRepository.FindOne(conditions, creatureFilters)
}
