package filters

import "gorm.io/gorm"

// OrderByFilter ORDER BY 排序
// value 格式："{欄位} ASC" 或 "{欄位} DESC"
// 範例："created_at DESC"、"name ASC"
type OrderByFilter struct {
	Value string
}

func (f OrderByFilter) Apply(db *gorm.DB) *gorm.DB {
	return db.Order(f.Value)
}

func NewOrderByFilter(value string) OrderByFilter {
	return OrderByFilter{Value: value}
}
