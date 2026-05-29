package seeder

import "game-backend/logger"

// SeedPlayers 寫入初始玩家資料（測試用）
// TODO: 待玩家 API 實作後補充測試用玩家資料
func SeedPlayers(log logger.Logger) {
	// 目前無需預建玩家，玩家由前端呼叫 API 自行建立
	log.Info("Seeder：玩家資料（略過，玩家由 API 建立）")
}
