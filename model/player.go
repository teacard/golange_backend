package model

import "gorm.io/gorm"

// Player 對應資料庫的 players 資料表
//
// 使用 gorm.Model（含 deleted_at）的原因：
//   → 玩家帳號可能被停用，需要軟刪除保留歷史紀錄
//
// json tag 不在這裡定義，API 回傳格式由 DTO 負責
type Player struct {
	gorm.Model

	// PlayerID 是業務層的玩家識別碼
	// 與資料庫主鍵 ID（uint）分開，讓外部 API 不暴露資料庫內部流水號
	PlayerID string
	Name     string
}
