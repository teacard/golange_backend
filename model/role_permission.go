package model

// RolePermission 對應資料庫的 role_permissions 資料表（pivot）
//
// 無 id、created_at、updated_at：純粹作為關聯 pivot，不需要額外欄位
// UNIQUE(role_id, permission_id) 由 migration 在資料庫層保證
type RolePermission struct {
	RoleID       uint `gorm:"primaryKey"`
	PermissionID uint `gorm:"primaryKey"`
}
