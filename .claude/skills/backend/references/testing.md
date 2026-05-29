# 測試規範

每支 API 都必須有測試，測試與實作放在同一目錄：

```
handler/player/
├── player.go
└── player_test.go
```

---

## 測試範本

```go
// handler/player/player_test.go
package player_test

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestCreatePlayer(t *testing.T) {
    // 1. 建立測試用 Gin engine（Test mode 關掉多餘 log）
    gin.SetMode(gin.TestMode)
    r := gin.New()
    r.POST("/api/players", Create)

    // 2. 建立假請求
    body := `{"name":"測試玩家"}`
    req := httptest.NewRequest(http.MethodPost, "/api/players", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    // 3. 驗證結果
    assert.Equal(t, http.StatusCreated, w.Code)

    var res PlayerResponse
    json.Unmarshal(w.Body.Bytes(), &res)
    assert.Equal(t, "測試玩家", res.Name)
    assert.NotEmpty(t, res.PlayerID)
}
```

---

## 每支 API 至少涵蓋的測試案例

| 案例 | 預期 HTTP 狀態碼 |
|------|----------------|
| 正常流程（正確輸入） | 200 / 201 |
| 必填欄位缺少 | 400 |
| 格式不符（長度超限、型別錯誤） | 400 |
| 資源不存在 | 404 |
| 重複建立（唯一值衝突） | 409 |

---

## 執行測試

```bash
# 執行全部測試
make test

# 執行單一套件
go test ./handler/player/...

# 顯示詳細輸出
go test -v ./handler/player/...

# 顯示覆蓋率
go test -cover ./...
```
