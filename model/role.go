package model

// Role 對應資料庫的 roles 資料表
//
// 使用 ChildModel（不含 deleted_at）的原因：
//   → 角色是系統靜態設定，不應被軟刪除
//
// roles 同時供 admins 與 users 使用，透過 model_roles 做多型關聯
type Role struct {
	ChildModel

	Name        string       // 角色名稱，例如：super_admin / editor
	Permissions []Permission `gorm:"many2many:role_permissions;"` // 多對多 → role_permissions
}
