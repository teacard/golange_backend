package router

import (
	_ "embed" // 啟用 //go:embed 指令，讓 Go 在編譯時把靜態檔案內容打包進執行檔
	"encoding/json"
	"net/http"

	"game-backend/config"
	docs "game-backend/docs" // swag init 產生的 swagger 規格資料

	swaggerFiles "github.com/swaggo/files"     // 內嵌的 Swagger UI 靜態資源（js、css 等）
	ginSwagger "github.com/swaggo/gin-swagger" // 將 Swagger UI 掛載到 Gin 的輔助套件

	"github.com/gin-gonic/gin"
)

// swaggerIndexHTML 在編譯時自動將 swagger_index.html 的內容嵌入這個字串變數。
// 這樣部署時不需要另外攜帶 html 檔案，執行檔本身就包含了這份 HTML。
//
//go:embed swagger_index.html
var swaggerIndexHTML string

// buildSwaggerSpec 以 swag 產生的規格為基礎，複製一份並覆寫 host 與 schemes，
// 回傳修改後的 JSON bytes。
//
// 這樣做的原因：swag 產生的規格 host 是固定的（localhost:8000），
// 但我們需要同時提供「本機」和「測試站」兩種規格供 Swagger UI 切換使用。
func buildSwaggerSpec(host string, schemes []string) ([]byte, error) {
	// docs.SwaggerInfo.ReadDoc() 會把 swag 產生的 JSON 樣板渲染成完整字串
	var spec map[string]any
	if err := json.Unmarshal([]byte(docs.SwaggerInfo.ReadDoc()), &spec); err != nil {
		return nil, err
	}
	spec["host"] = host
	spec["schemes"] = schemes
	return json.Marshal(spec)
}

// registerSwaggerRoutes 掛載所有 Swagger 相關路由：
//   - /swagger/spec-local.json  → 本機規格（host: localhost:8000）
//   - /swagger/spec-test.json   → 測試站規格（host: familiar.teacard-side-project.org）
//   - /swagger/admin-api/*      → Swagger UI 介面（後台 API 文件頁面）
func init() {
	registerGroup(registerSwaggerRoutes)
}

func registerSwaggerRoutes(r *gin.Engine) {
	// 本機開發用規格
	r.GET("/swagger/spec-local.json", func(c *gin.Context) {
		data, err := buildSwaggerSpec("localhost:8000", []string{"http"})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Data(http.StatusOK, "application/json; charset=utf-8", data)
	})

	// 測試站規格：僅在 .env 設定了 SWAGGER_TEST_HOST 時才掛載此路由
	if host := config.App.SwaggerTestHost; host != "" {
		r.GET("/swagger/spec-test.json", func(c *gin.Context) {
			data, err := buildSwaggerSpec(host, []string{"https"})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.Data(http.StatusOK, "application/json; charset=utf-8", data)
		})
	}

	// ginSwagger.WrapHandler 把 swaggo 內嵌的 Swagger UI 靜態資源包成 Gin handler，
	// 讓我們可以在 catch-all 路由（*any）中統一處理所有靜態檔案請求。
	handler := ginSwagger.WrapHandler(swaggerFiles.Handler)

	// *any 是 Gin 的 catch-all 參數，會匹配 /swagger/admin-api/ 後面的所有路徑。
	// 我們在這裡攔截特定路徑，其餘的交給 ginSwagger handler 處理。
	r.GET("/swagger/admin-api/*any", func(c *gin.Context) {
		switch c.Param("any") {
		case "/":
			// 根路徑直接導向 index.html，避免空路徑造成 404
			c.Redirect(http.StatusMovedPermanently, "/swagger/admin-api/index.html")
		case "/index.html":
			// 攔截 index.html，改回傳我們自訂的版本（內含 Base URL 切換選單）
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(http.StatusOK, swaggerIndexHTML)
		default:
			// 其他靜態資源（swagger-ui.css、swagger-ui-bundle.js 等）交給 ginSwagger 處理
			handler(c)
		}
	})
}
