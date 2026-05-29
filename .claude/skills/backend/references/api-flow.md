# API 完整流程

## 執行順序

```
Request
  │
  ▼
[Router]            ← 比對路由，找到對應的 handler chain（middleware + handler）
  │
  ▼
[Middleware]        ← 身份驗證（JWT）
  │                   未通過 → 直接回 401，不進入 Handler
  ▼
[Handler]           ← handler/{model}/{model}.go
  ├─ 1. Request DTO Bind & Validate  ← 檢查欄位格式（必填、長度、型別）
  │                                     格式不符 → 回 400
  ├─ 2. 呼叫 Service                 ← 傳入業務所需參數
  └─ 3. 回傳 Response DTO            ← 轉換成 API 輸出格式
          │
          ▼
      [Service]     ← service/{model}/{model}.go，業務邏輯
          │
          ▼
      [Repository]  ← repository/{model}/{model}.go，資料庫操作
          │
          ▼
          DB
```

**Gin 執行機制：**
Router 先比對路由，找到整條執行鏈（middleware + handler）後依序執行。
Middleware 一定在 Handler 之前——身份沒過不會跑到 Request 驗證。

---

## Router 設置範本

```go
// router/router.go
func Setup() *gin.Engine {
    r := gin.Default()
    api := r.Group("/api")

    // 不需要登入
    api.GET("/creatures",     creatureHandler.GetAll)
    api.GET("/creatures/:id", creatureHandler.GetByID)

    // 需要登入，掛 AuthMiddleware
    authed := api.Group("/", middleware.AuthMiddleware())
    {
        authed.GET("/players/me",           playerHandler.GetMe)
        authed.GET("/players/me/inventory", playerHandler.GetInventory)
        authed.POST("/players/:id/catch",   catchHandler.Catch)
    }

    return r
}
```

---

## Handler 範本

```go
// handler/player/player.go
func Create(c *gin.Context) {
    // 1. Bind & Validate
    var req CreatePlayerRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 2. 呼叫 Service
    player, err := playerService.Create(req.Name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 3. 回傳 Response DTO
    c.JSON(http.StatusCreated, toPlayerResponse(player))
}
```

---

## Middleware 取得登入玩家

Middleware 驗證成功後把 `player_id` 寫進 Gin context，Handler 直接取用：

```go
// handler 內取得目前登入的玩家 ID
playerID := c.GetUint("player_id")
```
