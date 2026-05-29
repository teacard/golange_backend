package filters

import "gorm.io/gorm"

// CompareFilter 對應 GT / GTE / LT / LTE
// 用 constructor function 建立，不直接用 struct
type CompareFilter struct {
	Column   string
	Operator string
	Value    interface{}
}

func (f CompareFilter) Apply(db *gorm.DB) *gorm.DB {
	return db.Where(f.Column+" "+f.Operator+" ?", f.Value)
}

// GT WHERE column > value
func GT(column string, value interface{}) CompareFilter {
	return CompareFilter{Column: column, Operator: ">", Value: value}
}

// GTE WHERE column >= value
func GTE(column string, value interface{}) CompareFilter {
	return CompareFilter{Column: column, Operator: ">=", Value: value}
}

// LT WHERE column < value
func LT(column string, value interface{}) CompareFilter {
	return CompareFilter{Column: column, Operator: "<", Value: value}
}

// LTE WHERE column <= value
func LTE(column string, value interface{}) CompareFilter {
	return CompareFilter{Column: column, Operator: "<=", Value: value}
}
