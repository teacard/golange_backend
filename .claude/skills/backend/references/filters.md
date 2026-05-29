# Filter 系統說明

## 呼叫方式（service 層）

```go
// 通用 filters（CommonFilters 內建）
creatures, err := repo.FindAll(map[string]any{
    "id":       []int{1, 2, 3},    // WHERE id IN (1,2,3)
    "order_by": "created_at DESC", // ORDER BY created_at DESC
    "limit":    10,                // LIMIT 10
    "offset":   0,                 // OFFSET 0
})

// model 專屬 filters
creatures, err := repo.FindAll(map[string]any{
    "rarity":  "uncommon",
    "habitat": "草原",
})
```

---

## 解析順序

```
map key
  │
  ▼
CommonFilters（repository/filter_factory.go）
  │ 找不到
  ▼
model 專屬 FilterFactory（repository/{model}/{model}.go 內的 {model}Filters）
  │ 找不到
  ▼
error：filter key 'xxx' 未被註冊
```

---

## 通用 filter key 清單

| key | value 型別 | 產生的 SQL |
|-----|-----------|-----------|
| `id` | `int` 或 `[]int` | `WHERE id = ?` 或 `WHERE id IN (?)` |
| `order_by` | `string` | `ORDER BY {value}` |
| `limit` | `int` | `LIMIT {value}` |
| `offset` | `int` | `OFFSET {value}` |

---

## 新增 model 專屬 filter

**Step 1**：在 `repository/{model}/filters/` 建立新檔案

```go
// repository/creature/filters/personality.go
package creaturefilters

import (
    "fmt"
    "game-backend/repository"
    "game-backend/repository/filters"
)

func PersonalityFactory(v any) (repository.FilterInterface, error) {
    s, ok := v.(string)
    if !ok {
        return nil, fmt.Errorf("personality filter: value 必須是 string")
    }
    return filters.NewEqFilter("personality", s), nil
}
```

**Step 2**：在 `repository/{model}/{model}.go` 的 `{model}Filters` 登記

```go
var creatureFilters = repository.FilterFactory{
    "rarity":      creaturefilters.RarityFactory,
    "habitat":     creaturefilters.HabitatFactory,
    "personality": creaturefilters.PersonalityFactory, // ← 新增
}
```

---

## 目前已註冊的 model 專屬 filter

### CreatureFilters
| key | 說明 |
|-----|------|
| `rarity` | 稀有度，支援 string 或 []string |
| `habitat` | 棲息地，string |

### PlayerFilters
| key | 說明 |
|-----|------|
| `name` | 玩家名稱，string |
