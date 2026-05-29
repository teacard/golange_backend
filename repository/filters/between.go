package filters

import "gorm.io/gorm"

// BetweenFilter WHERE column BETWEEN min AND max
// 用法：filters.NewBetweenFilter("created_at", startTime, endTime)
type BetweenFilter struct {
	Column string
	Min    interface{}
	Max    interface{}
}

func (f BetweenFilter) Apply(db *gorm.DB) *gorm.DB {
	return db.Where(f.Column+" BETWEEN ? AND ?", f.Min, f.Max)
}

func NewBetweenFilter(column string, min, max interface{}) BetweenFilter {
	return BetweenFilter{Column: column, Min: min, Max: max}
}
