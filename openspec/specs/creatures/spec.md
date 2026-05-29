# Creatures Specification

## Purpose

管理遊戲中所有可收集生物的資料，包含外觀、動畫參數與多語系名稱。

---

## Requirements

### Requirement: List All Creatures

系統 SHALL 提供端點讓前端取得所有生物的完整資料清單。

#### Scenario: 成功取得生物列表

- WHEN 前端發送 `GET /api/creatures`
- THEN 系統回傳 HTTP 200
- AND 回傳所有生物資料，包含 id、rarity、personality、habitat、colors、special_parts、animation

#### Scenario: 無生物資料時

- WHEN 資料庫中無生物資料
- THEN 系統回傳 HTTP 200
- AND data 為空陣列

---

### Requirement: Get Single Creature

系統 SHALL 提供端點讓前端依 creature_id 取得單一生物資料。

#### Scenario: 成功取得指定生物

- WHEN 前端發送 `GET /api/creatures/:id`，且該 creature_id 存在
- THEN 系統回傳 HTTP 200
- AND 回傳該生物完整資料

#### Scenario: 生物不存在

- WHEN 前端發送 `GET /api/creatures/:id`，且該 creature_id 不存在
- THEN 系統回傳 HTTP 404
- AND 回傳 `{"error": "找不到該生物"}`

---

### Requirement: Creature Data Structure

生物資料 SHALL 包含以下欄位：

- `id`：唯一識別碼，格式 `creature_<英文名稱小寫>`
- `rarity`：稀有度，可選值 common / uncommon / rare / legendary
- `personality`：個性類型
- `habitat`：棲息地（用於篩選）
- `colors`：JSONB，格式 `{"primary":"#RRGGBB","secondary":"#RRGGBB"}`
- `special_parts`：JSONB，格式 `[{"code":"string","color":"#RRGGBB（選填）"}]`
- `animation`：JSONB，依狀態分組 `{"idle":{...},"walk":{...},"catch":{...}}`

#### Scenario: 多語系名稱存於獨立資料表

- WHEN 系統儲存生物資料
- THEN 名稱與說明文字存於 creature_locales
- AND 以 (creature_id, lang_code) 作為唯一索引
