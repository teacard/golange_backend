# DTO 設計規範

DTO（Data Transfer Object）分兩種，都放在 `handler/{model}/` 內。

---

## Request DTO（輸入驗證）

負責定義前端傳入的格式，並用 `binding` tag 做驗證。

```go
type CreatePlayerRequest struct {
    // binding:"required"      → 必填，空字串視為未填
    // binding:"min=1,max=50"  → 長度限制 1~50 字元
    Name string `json:"name" binding:"required,min=1,max=50"`
}
```

常用 binding 規則：

| tag | 說明 |
|-----|------|
| `required` | 必填 |
| `min=N` | 最小長度（字串）或最小值（數字） |
| `max=N` | 最大長度或最大值 |
| `email` | 必須是 email 格式 |
| `oneof=a b c` | 只允許指定值，例如 `oneof=common uncommon rare` |

Gin 透過 `ShouldBindJSON` 自動驗證，失敗直接回 400：

```go
var req CreatePlayerRequest
if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
}
```

---

## Response DTO（輸出格式）

負責定義 API 回傳的 JSON 欄位，只暴露需要的資訊。

```go
type PlayerResponse struct {
    PlayerID  string `json:"player_id"`
    Name      string `json:"name"`
    CreatedAt string `json:"created_at"`
}

// 從 model 轉換成 DTO 的轉換函式
func toPlayerResponse(p model.Player) PlayerResponse {
    return PlayerResponse{
        PlayerID:  p.PlayerID,
        Name:      p.Name,
        CreatedAt: p.CreatedAt.Format(time.RFC3339),
    }
}
```

---

## 為什麼 Model 不放 json tag？

Model 職責是對應資料庫，json 輸出格式是 API 的責任。

分開的好處：
- 同一個 model 可以對應多個 DTO（詳細版、列表精簡版、內嵌版）
- 調整 API 輸出格式不需要動 model
- 避免直接把 `deleted_at`、DB 內部 ID 等資料庫欄位暴露給前端

```go
// ✅ model 不放 json tag
type Player struct {
    gorm.Model
    PlayerID string
    Name     string
}

// ✅ json tag 只在 DTO
type PlayerResponse struct {
    PlayerID string `json:"player_id"`
    Name     string `json:"name"`
}
```
