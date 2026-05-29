package model

import "time"

// ChildModel 子表專用的基底結構
//
// 什麼是子表？
//   → 內容依附於父表存在的附屬表，例如 creature_locales 依附於 creatures
//   → 父表軟刪除後，子表資料透過 JOIN 父表的 WHERE deleted_at IS NULL 自動過濾
//   → 子表本身不需要獨立的 deleted_at 欄位
//
// 為什麼不直接用 gorm.Model？
//   → gorm.Model 內含 DeletedAt 欄位，GORM 查詢時會自動加 WHERE deleted_at IS NULL
//   → 但子表的 SQL 沒有 deleted_at 欄位，這樣查詢會報錯
//   → 所以子表用這個自訂結構，只保留 id / created_at / updated_at
type ChildModel struct {
	// gorm:"primarykey" 告訴 GORM 這個欄位是主鍵
	// 這是查詢、更新、刪除時 GORM 定位資料列的依據，必須保留
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
