package catch

// CatchCreatureFlow 捕獲生物流程
// 跨多個 service 的複雜情境，例如：
// 1. 確認生物存在
// 2. 確認玩家尚未擁有該生物
// 3. 寫入玩家背包
// 4. 回傳捕獲結果

// TODO: 實作時注入所需的 service
// type CatchCreatureFlow struct {
//     creatureService *creature.Service
//     playerService   *player.Service
// }
//
// func (f *CatchCreatureFlow) Execute(playerID uint, creatureID string) error { ... }
