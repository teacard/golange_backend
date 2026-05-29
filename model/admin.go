package model

import (
	"time"

	"gorm.io/gorm"
)

// Admin 對應資料庫的 admins 資料表
//
// 使用 gorm.Model（含 deleted_at）的原因：
//   → 管理員帳號停用時需要軟刪除，保留操作記錄
//
// 與 roles 的關係：多對多（透過 model_roles 多型 pivot），實際使用上一對一
type Admin struct {
	gorm.Model

	Email       string     // 登入帳號，唯一
	Password    string     // bcrypt 雜湊後的密碼
	Name        string     // 顯示名稱
	Status      string     // 帳號狀態：active / suspended
	LastLoginAt *time.Time // 最後登入時間，nullable
}
