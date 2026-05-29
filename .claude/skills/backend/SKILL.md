---
name: backend
description: >
  草原精靈獸遊戲 — 後端程式開發規範。
  新增 API、model、filter、repository、service 時必須參考此規範。
---

# 後端開發規範

## 架構分層

| 層 | 職責 |
|----|------|
| `handler/` | HTTP 進出口：Bind Request DTO → 呼叫 Service → 回傳 Response DTO |
| `flow/` | 跨多個 service 的複雜情境（如捕獲生物） |
| `service/` | 單一資源的商業邏輯 |
| `repository/` | 資料庫查詢抽象（map-based filter） |
| `model/` | 對應資料庫資料表結構 |
| `seeder/` | 初始資料填充 |
| `migrations/` | 版本化 SQL 建表（見 migration skill） |

---

## 詳細說明（按需載入）

| 需要做什麼 | 讀哪個檔案 |
|-----------|-----------|
| 了解 API 完整流程、Router/Middleware/Handler 範本 | `.claude/skills/backend/references/api-flow.md` |
| Request DTO 驗證、Response DTO 格式設計 | `.claude/skills/backend/references/dto.md` |
| Filter map 用法、新增專屬 filter | `.claude/skills/backend/references/filters.md` |
| Model 主表/子表設計、FK 規範 | `.claude/skills/backend/references/models.md` |
| 測試範本、測試案例覆蓋要求 | `.claude/skills/backend/references/testing.md` |

---

## 新增 API Checklist

- [ ] `migrations/` 確認資料表 SQL（參考 migration skill）
- [ ] `model/` 確認 struct 欄位
- [ ] `repository/{model}/filters/` 新增需要的 filter
- [ ] `service/{model}/` 實作業務邏輯
- [ ] `handler/{model}/` 定義 Request DTO + Response DTO + Handler
- [ ] `handler/{model}/` 寫測試（正常流程 + 異常案例）
- [ ] `router/router.go` 登記路由（確認是否需要 middleware）
- [ ] `make test` 全部通過

---

## 命名規範

| 類型 | 規範 | 範例 |
|------|------|------|
| Filter key | snake_case | `order_by`、`rarity` |
| Go 變數/函式 | camelCase | `findAll`、`playerID` |
| Go 型別/struct | PascalCase | `PlayerRepository`、`CreatePlayerRequest` |
| 資料庫欄位 | snake_case（GORM 自動轉換） | `created_at`、`player_id` |
| API 路由 | snake_case | `/api/players/:player_id` |
| Business ID | `{資源}_{英文名}` | `creature_grassspirit` |
