# 資料庫設計

## creatures

| 欄位 | 類型 | 說明 |
|------|------|------|
| id | PK | 自動遞增 |
| creature_id | string UNIQUE | 唯一識別碼，格式 `creature_<英文名稱小寫>` |
| rarity | string | 稀有度：common / uncommon / rare / legendary |
| personality | string | 個性類型，例如：悠閒型、活潑型 |
| habitat | string | 棲息地，例如：草原、森林、海洋 |
| colors | JSONB | 顏色配置 |
| special_parts | JSONB | 特殊部位清單（僅供前端繪製） |
| animation | JSONB | 動畫參數（依狀態分組） |

**JSONB 格式**

```json
// colors
{"primary": "#e8d8b0", "secondary": "#b8a070"}

// special_parts — code 必填，color 選填（覆蓋預設色）
[
  {"code": "head_leaf", "color": "#8aba60"},
  {"code": "round_ears"}
]

// animation
{
  "idle":  {"breath_speed": 1.2, "breath_amp": 3, "tail_speed": 0.9},
  "walk":  {"cycle": 3.0, "leg_amp": 0.35, "body_bounce": 5},
  "catch": {"shake_freq": 18, "duration": 1.5, "blush_alpha": 0.6}
}
```

---

## creature_locales

| 欄位 | 類型 | 說明 |
|------|------|------|
| id | PK | 自動遞增 |
| creature_id | string | FK → creatures.creature_id |
| lang_code | string | 語言代碼：zh-TW / en / ja / ko… |
| name | string | 該語言的生物名稱 |
| description | string | 該語言的圖鑑說明文字 |
| UNIQUE | — | (creature_id, lang_code) 組合唯一 |

---

## players

| 欄位 | 類型 | 說明 |
|------|------|------|
| id | PK | 自動遞增 |
| player_id | string UNIQUE | 玩家唯一識別碼 |
| name | string | 玩家顯示名稱 |

---

## inventories

| 欄位 | 類型 | 說明 |
|------|------|------|
| id | PK | 自動遞增 |
| player_id | uint | FK → players.id |
| creature_id | string | FK → creatures.creature_id |
| caught_at | string | 捕獲時間（ISO 8601） |

---

## 設計原則

- `rarity`、`personality`、`habitat` 為一般欄位，用於查詢篩選
- `colors`、`special_parts`、`animation` 為 JSONB，擴充不需改欄位
- 多語系資料統一放 `creature_locales`，主表不存任何語系文字
- 所有欄位皆有 GORM comment，PostgreSQL 可直接查看用途
