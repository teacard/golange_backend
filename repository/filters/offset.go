package filters

import "gorm.io/gorm"

// OffsetFilter OFFSET 跳過前 N 筆
// 搭配 LimitFilter 做分頁
// 範例：limit=10, offset=20 → 第 3 頁（跳過前 20 筆，取第 21~30 筆）
type OffsetFilter struct {
	Value int
}

func (f OffsetFilter) Apply(db *gorm.DB) *gorm.DB {
	return db.Offset(f.Value)
}

func NewOffsetFilter(value int) OffsetFilter {
	return OffsetFilter{Value: value}
}
