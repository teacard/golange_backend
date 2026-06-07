package model

import permission "game-backend/enum/permission"

// Permission 對應資料庫的 permissions 資料表
//
// 使用 ChildModel（不含 deleted_at）的原因：
//   → 權限是系統靜態設定，不應被軟刪除，異動直接修改資料
type Permission struct {
	ChildModel

	Name string          // 顯示名稱，例如：後台人員-檢視
	Code permission.Code // 程式識別碼，使用 enum/permission 常數
}
