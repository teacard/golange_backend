package filters

import "gorm.io/gorm"

// EqFilter WHERE column = value
// 用法：filters.NewEqFilter("rarity", "uncommon")
type EqFilter struct {
	Column string
	Value  interface{}
}

func (f EqFilter) Apply(db *gorm.DB) *gorm.DB {
	return db.Where(f.Column+" = ?", f.Value)
}

func NewEqFilter(column string, value interface{}) EqFilter {
	return EqFilter{Column: column, Value: value}
}
