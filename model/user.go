package model

import (
	"time"

	"gorm.io/gorm"
)

// User 對應資料庫的 users 資料表
//
// 使用 gorm.Model（含 deleted_at）的原因：
//   → 玩家帳號停權時需要軟刪除，保留遊戲資料
//
// UserID 使用 ULID（26 字元）作為對外識別碼，不暴露資料庫流水號
// 生成邏輯在 service 層的 BeforeCreate hook 中處理
type User struct {
	gorm.Model

	UserID      string     // ULID，對外識別碼，唯一，26 字元
	Email       string     // 登入帳號，唯一
	Password    string     // bcrypt 雜湊後的密碼
	Nickname    string     // 遊戲內顯示名稱
	Status      string     // 帳號狀態：active / suspended
	LastLoginAt *time.Time // 最後登入時間，nullable
}
