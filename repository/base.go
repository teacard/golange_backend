package repository

import (
	"maps"

	"gorm.io/gorm"
)

// BaseRepository 泛型通用 Repository
// T 可以是任何 model（Creature、Player 等）
//
// [T any] 是 Go 的泛型語法，代表「T 可以是任何型別」
// 這樣同一份程式碼就能處理不同的 model，不需要為每個 model 重複寫相同的 CRUD
type BaseRepository[T any] struct {
	DB *gorm.DB
}

// resolveConditions 將 map[string]any 條件解析成 []FilterInterface
// 解析順序：CommonFilters → extraFactories
// 任何一個 key 找不到就回傳 error，讓呼叫端明確知道哪個 key 未被註冊
func (r *BaseRepository[T]) resolveConditions(conditions map[string]any, extraFactories []FilterFactory) ([]FilterInterface, error) {
	result := make([]FilterInterface, 0, len(conditions))
	for key, value := range conditions {
		f, err := Resolve(key, value, extraFactories...)
		if err != nil {
			return nil, err
		}
		result = append(result, f)
	}
	return result, nil
}

// applyFilters 依序把所有 filter 套用到 GORM query 上
// GORM 的 *gorm.DB 採用 chain 設計，每次 Where/Order/Limit 都回傳新物件，不修改原本的
func (r *BaseRepository[T]) applyFilters(fs []FilterInterface) *gorm.DB {
	query := r.DB
	for _, f := range fs {
		query = f.Apply(query)
	}
	return query
}

// FindAll 查詢多筆資料
//
// conditions：filter 條件 map，key 為 filter 名稱（snake_case），value 為條件值
// extraFactories：model 專屬的 FilterFactory，由子 repository 的 FindAll 傳入
//
// 使用範例（從 service 層透過 CreatureRepository 呼叫）：
//
//	creatures, err := repo.FindAll(map[string]any{
//	    "rarity":   "uncommon",
//	    "order_by": "created_at DESC",
//	    "limit":    10,
//	})
func (r *BaseRepository[T]) FindAll(conditions map[string]any, extraFactories ...FilterFactory) ([]T, error) {
	fs, err := r.resolveConditions(conditions, extraFactories)
	if err != nil {
		return nil, err
	}
	var results []T
	err = r.applyFilters(fs).Find(&results).Error
	return results, err
}

// FindOne 查詢單筆（第一筆符合條件的資料）
// 找不到時 GORM 會回傳 gorm.ErrRecordNotFound
func (r *BaseRepository[T]) FindOne(conditions map[string]any, extraFactories ...FilterFactory) (T, error) {
	fs, err := r.resolveConditions(conditions, extraFactories)
	var result T
	if err != nil {
		return result, err
	}
	err = r.applyFilters(fs).First(&result).Error
	return result, err
}

// Exists 確認是否存在符合條件的資料，回傳 true/false
func (r *BaseRepository[T]) Exists(conditions map[string]any, extraFactories ...FilterFactory) (bool, error) {
	fs, err := r.resolveConditions(conditions, extraFactories)
	if err != nil {
		return false, err
	}
	var count int64
	var zero T
	err = r.applyFilters(fs).Model(&zero).Count(&count).Error
	return count > 0, err
}

// Count 計算符合條件的資料筆數
func (r *BaseRepository[T]) Count(conditions map[string]any, extraFactories ...FilterFactory) (int64, error) {
	fs, err := r.resolveConditions(conditions, extraFactories)
	if err != nil {
		return 0, err
	}
	var count int64
	var zero T
	err = r.applyFilters(fs).Model(&zero).Count(&count).Error
	return count, err
}

// PaginationResult 分頁查詢結果，包含資料列表與總筆數
type PaginationResult[T any] struct {
	Data  []T
	Total int64
}

// Paginate 執行分頁查詢，內部自動合併 FindAll 與 Count 兩次查詢
// conditions 只需帶過濾條件，limit/offset 由 page、perPage 自動計算
func (r *BaseRepository[T]) Paginate(page, perPage int, conditions map[string]any, extraFactories ...FilterFactory) (PaginationResult[T], error) {
	total, err := r.Count(conditions, extraFactories...)
	if err != nil {
		return PaginationResult[T]{}, err
	}

	dataConditions := map[string]any{
		"limit":  perPage,
		"offset": (page - 1) * perPage,
	}
	maps.Copy(dataConditions, conditions)

	data, err := r.FindAll(dataConditions, extraFactories...)
	if err != nil {
		return PaginationResult[T]{}, err
	}

	return PaginationResult[T]{Data: data, Total: total}, nil
}

// Create 新增一筆資料
// entity 傳指標（*T），GORM INSERT 後會把產生的 ID 寫回 entity
func (r *BaseRepository[T]) Create(entity *T) error {
	return r.DB.Create(entity).Error
}

// Save 更新一筆資料（有 PK 就 UPDATE，沒有就 INSERT）
func (r *BaseRepository[T]) Save(entity *T) error {
	return r.DB.Save(entity).Error
}

// Delete 軟刪除一筆資料（填入 deleted_at，不真正刪除資料列）
func (r *BaseRepository[T]) Delete(entity *T) error {
	return r.DB.Delete(entity).Error
}
