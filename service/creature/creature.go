package creature

import (
	"game-backend/db"
	"game-backend/model"
)

// GetAll 取得所有生物
func GetAll() ([]model.Creature, error) {
	var creatures []model.Creature
	result := db.DB.Find(&creatures)
	return creatures, result.Error
}

// GetByID 取得單一生物
func GetByID(creatureID string) (model.Creature, error) {
	var creature model.Creature
	result := db.DB.Where("creature_id = ?", creatureID).First(&creature)
	return creature, result.Error
}
