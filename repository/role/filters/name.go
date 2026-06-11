package rolefilters

import (
	"game-backend/repository"
	"gorm.io/gorm"
)

// iLikeFilter WHERE name ILIKE '%value%'（不區分大小寫模糊搜尋）
type iLikeFilter struct {
	value string
}

func (f iLikeFilter) Apply(db *gorm.DB) *gorm.DB {
	return db.Where("name ILIKE ?", "%"+f.value+"%")
}

// NameFactory 角色名稱 filter（不區分大小寫模糊比對）
// value 由 service 層傳入，呼叫端保證型別為 string
func NameFactory(raw any) (repository.FilterInterface, error) {
	return iLikeFilter{value: raw.(string)}, nil
}
