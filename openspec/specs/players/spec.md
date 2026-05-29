# Players Specification

## Purpose

管理玩家資料與玩家收集的生物背包。

---

## Requirements

### Requirement: Player Data Structure

玩家資料 SHALL 包含以下欄位：

- `player_id`：玩家唯一識別碼
- `name`：玩家顯示名稱

#### Scenario: 玩家識別碼唯一

- WHEN 建立新玩家
- THEN player_id 不得與現有玩家重複
- AND 系統回傳錯誤訊息若重複

---

### Requirement: Player Inventory（待實作）

系統 SHALL 提供端點讓前端取得玩家已收集的生物清單。

#### Scenario: 取得玩家背包

- WHEN 前端發送 `GET /api/player/inventory`
- THEN 系統回傳該玩家所有已收集生物的 creature_id 清單與捕獲時間

---

### Requirement: Catch Creature（待實作）

系統 SHALL 提供端點讓前端觸發捕獲生物流程。

#### Scenario: 成功捕獲

- WHEN 前端發送 `POST /api/player/catch`，帶有 player_id 與 creature_id
- AND 該生物存在
- AND 玩家尚未擁有該生物
- THEN 系統將生物寫入玩家背包
- AND 回傳 HTTP 200 與捕獲結果

#### Scenario: 生物不存在

- WHEN 前端發送 `POST /api/player/catch`，但 creature_id 不存在
- THEN 系統回傳 HTTP 404

#### Scenario: 玩家已擁有該生物

- WHEN 前端發送 `POST /api/player/catch`，但玩家已擁有該生物
- THEN 系統回傳 HTTP 409
