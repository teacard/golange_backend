package filters

// IdFilter 對應 PHP 的 Id 繼承 In
// 所有資料表的主鍵查詢都用這個，不需要每個 repository 自己重寫
// 用法：filters.NewIdFilter(1, 2, 3)
type IdFilter struct {
	InFilter
}

func NewIdFilter(ids ...interface{}) IdFilter {
	return IdFilter{InFilter{Column: "id", Values: ids}}
}
