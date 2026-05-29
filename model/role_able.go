package model

// ModelRole 對應資料庫的 model_roles 資料表（polymorphic pivot）
//
// admins 與 users 共用此表，透過 Model 欄位區分來源：
//   - Model = "admin" → ModelID 對應 admins.id
//   - Model = "user"  → ModelID 對應 users.id
//
// 無 id、created_at、updated_at：純粹作為關聯 pivot
// UNIQUE(model, model_id, role_id) 由 migration 在資料庫層保證
type ModelRole struct {
	Model   string `gorm:"primaryKey"` // 來源 model：admin / user
	ModelID uint   `gorm:"primaryKey"` // 對應各自資料表的 id
	RoleID  uint   `gorm:"primaryKey"` // FK → roles.id
}

// ModelRoleModel 定義 model 欄位的合法值（enum）
type ModelRoleModel string

const (
	ModelRoleModelAdmin ModelRoleModel = "admin"
	ModelRoleModelUser  ModelRoleModel = "user"
)
