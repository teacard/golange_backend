package filters

import "gorm.io/gorm"

// LimitFilter LIMIT 限制回傳筆數
// 通常搭配 OffsetFilter 做分頁
// 範例："limit": 10 → LIMIT 10
type LimitFilter struct {
	Value int
}

func (f LimitFilter) Apply(db *gorm.DB) *gorm.DB {
	return db.Limit(f.Value)
}

func NewLimitFilter(value int) LimitFilter {
	return LimitFilter{Value: value}
}
