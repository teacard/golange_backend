# Model 設計原則

## 主表用 `gorm.Model`（含 deleted_at）

適用：可以被獨立軟刪除的資料表（creatures、players）

```go
type Player struct {
    gorm.Model  // 提供 ID / CreatedAt / UpdatedAt / DeletedAt
    PlayerID string
    Name     string
}
```

`gorm.Model` 提供的欄位：
| 欄位 | 型別 | 說明 |
|------|------|------|
| `ID` | `uint` | 主鍵，自動遞增 |
| `CreatedAt` | `time.Time` | INSERT 時自動填入 |
| `UpdatedAt` | `time.Time` | UPDATE 時自動更新 |
| `DeletedAt` | `gorm.DeletedAt` | 軟刪除，Delete() 時填入，查詢自動加 WHERE deleted_at IS NULL |

---

## 子表用 `ChildModel`（無 deleted_at）

適用：內容依附父表存在，不需要獨立軟刪除（例如 creature_locales）

```go
type CreatureLocale struct {
    ChildModel   // 提供 ID / CreatedAt / UpdatedAt（無 DeletedAt）
    CreatureID uint
    LangCode   string
    Name       string
    Description string
}
```

為什麼不用 `gorm.Model`？
→ `gorm.Model` 有 `DeletedAt`，GORM 查詢會自動加 `WHERE deleted_at IS NULL`
→ 子表的 SQL 沒有 `deleted_at` 欄位，查詢會報錯
→ 子表可見性由 JOIN 父表的 `WHERE deleted_at IS NULL` 控制

---

## FK 統一用 DB primary key（uint）

```go
// ✅ 正確：FK 指向 DB primary key
PlayerID   uint  // → players.id
CreatureID uint  // → creatures.id

// ❌ 錯誤：FK 指向 business id（varchar）
PlayerID   string  // → players.player_id（不做 FK）
```

business id（`player_id varchar`、`creature_id varchar`）只用於：
- API 對外識別（URL、JSON response）
- Seeder 的 FirstOrCreate 條件

---

## Model 不放 json tag

json tag 由 DTO 負責。詳見 `.claude/skills/backend/references/dto.md`。

---

## 欄位命名對應

GORM 自動將 Go 的 CamelCase 轉成 SQL 的 snake_case：

| Go 欄位名 | SQL 欄位名 |
|-----------|-----------|
| `CreatureID` | `creature_id` |
| `PlayerID` | `player_id` |
| `SpecialParts` | `special_parts` |
| `CaughtAt` | `caught_at` |

不需要額外寫 `gorm:"column:xxx"` tag（除非命名有特殊需求）。
