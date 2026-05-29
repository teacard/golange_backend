package filters

import "gorm.io/gorm"

// InFilter 對應 PHP 的 In 抽象類
// 用法：InFilter{Column: "rarity", Values: []interface{}{"common", "uncommon"}}
type InFilter struct {
	Column string
	Values interface{}
}

func NewInFilter(column string, values interface{}) InFilter {
	return InFilter{Column: column, Values: values}
}

func (f InFilter) Apply(db *gorm.DB) *gorm.DB {
	return db.Where(f.Column+" IN ?", f.Values)
}
