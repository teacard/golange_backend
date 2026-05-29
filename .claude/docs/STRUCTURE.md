# 程式架構說明

## 請求流程

```
handler → flow（複雜跨資源） → service（共用邏輯） → repository → db
```
簡單請求可跳過 flow：`handler → service → repository → db`

---

## 各層職責

| 層 | 職責 | PHP 對應 |
|----|------|---------|
| `handler/` | 接收 HTTP 請求、回傳 JSON | Controller |
| `flow/` | 跨多個 service 的複雜情境 | Use Case |
| `service/` | 單一資源的共用商業邏輯 | Service |
| `repository/` | 資料庫操作抽象 | Repository |
| `model/` | 資料表結構定義 | Model |
| `migrations/` | SQL 版本化建表 / 回滾腳本 | Migration |
| `seeder/` | 初始資料填充 | Seeder |

---

## 目錄結構

```
backend/
├── main.go
├── config/config.go         ← .env 讀取，config.App.XXX 存取
├── logger/logger.go         ← Logger interface + logrus，測試可注入 mock
├── db/db.go                 ← PostgreSQL 連線 + 執行 golang-migrate
├── migrations/
│   ├── 000001_init_schema.up.sql   ← 建表 SQL
│   └── 000001_init_schema.down.sql ← 回滾 SQL
├── seeder/
│   ├── seeder.go            ← RunAll()：依序執行所有 seeder
│   ├── creature.go          ← 生物初始資料（從圖鑑讀取參數）
│   └── player.go            ← 玩家初始資料（目前為 placeholder）
├── model/
│   ├── creature.go          ← Creature + CreatureLocale
│   └── player.go            ← Player + Inventory
├── repository/
│   ├── filter.go            ← FilterInterface
│   ├── base.go              ← BaseRepository[T]（泛型）
│   ├── filters/
│   │   ├── in.go            ← InFilter
│   │   ├── id.go            ← IdFilter（內嵌 InFilter）
│   │   ├── eq.go            ← EqFilter
│   │   ├── between.go       ← BetweenFilter
│   │   └── compare.go       ← GT / GTE / LT / LTE
│   ├── creature/creature.go ← 生物專屬查詢
│   └── player/player.go     ← 玩家專屬查詢
├── router/router.go         ← 所有路由集中定義
├── handler/
│   ├── creature/creature.go + creature_test.go
│   └── player/player.go + player_test.go
├── service/
│   ├── creature/creature.go + creature_test.go
│   └── player/player.go + player_test.go
└── flow/
    └── catch/catch.go       ← 捕獲生物流程
```

---

## Migration 管理（golang-migrate）

命名規則：`{版本號}_{說明}.up.sql` / `{版本號}_{說明}.down.sql`

```
migrations/
├── 000001_init_schema.up.sql    ← 建立所有初始資料表
├── 000001_init_schema.down.sql  ← 刪除所有初始資料表
├── 000002_add_level.up.sql      ← （未來）新增欄位
└── 000002_add_level.down.sql    ← （未來）移除欄位
```

- **往前**：`db.Init()` 啟動時自動執行 `migrate.Up()`
- **回滾**：尚未提供 CLI 指令，未來可加入 `go run cmd/migrate/main.go down`
- **版本記錄**：postgresql 內的 `schema_migrations` 資料表

---

## Seeder 使用方式

```go
// main.go 啟動順序
db.Init(log)       // 1. 連線 + migration
seeder.RunAll(log) // 2. 填入初始資料
router.Setup()     // 3. 啟動 HTTP
```

新增生物 seeder 步驟：
1. 在 `.claude/creatures/` 建立圖鑑文件
2. 在 `seeder/creature.go` 新增 `seedXxx()` 函式
3. 在 `SeedCreatures()` 呼叫該函式

---

## Repository 使用方式

```go
repo := creature.NewCreatureRepository(db.DB)

// 查全部
all, _ := repo.FindAll()

// 帶 filter
uncommon, _ := repo.FindAll(filters.NewEqFilter("rarity", "uncommon"))

// 複合條件（自動 AND）
result, _ := repo.FindAll(
    filters.NewEqFilter("habitat", "草原"),
    filters.GTE("level", 3),
)

// 確認存在 / 計算筆數
exists, _ := repo.Exists(filters.NewEqFilter("creature_id", "creature_grassspirit"))
count,  _ := repo.Count(filters.NewEqFilter("rarity", "rare"))
```

---

## 注意事項

- 新增 API 前先在 `router/router.go` 登記路由
- 新增資料表用 SQL migration 檔案，不要直接改 AutoMigrate
- Logger 透過參數注入，不使用全域變數
- 種子資料寫在 `seeder/`，不在 `service/` 內
