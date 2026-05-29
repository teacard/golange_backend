package model

// Permission 對應資料庫的 permissions 資料表
//
// 使用 ChildModel（不含 deleted_at）的原因：
//   → 權限是系統靜態設定，不應被軟刪除，異動直接修改資料
//
// code 欄位作為程式內 enum 的對應依據，例如：UPLOAD_CREATURE
type Permission struct {
	ChildModel

	Name string // 顯示名稱，例如：上傳圖鑑
	Code string // 程式識別碼（enum），例如：UPLOAD_CREATURE
}
