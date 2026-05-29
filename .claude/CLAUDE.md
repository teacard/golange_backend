# 草原精靈獸遊戲 — 後端開發規範

## 技術棧

- 語言：Go
- 框架：Gin + GORM
- 資料庫：PostgreSQL 16（Docker）
- Migration：golang-migrate
- Logger：Logrus（interface 包裝）
- 設定：godotenv（.env 檔）

## 技能文件

- 後端開發規範 → `.claude/skills/backend/SKILL.md`
- Migration 規範 → `.claude/skills/migration/SKILL.md`

## 生物圖鑑

所有生物資料維護於 `.claude/creatures/` 目錄。
新增生物 seeder 前必須先建立圖鑑文件。

## 框架地圖

```
backend/
├── main.go              ← 進入點
├── config/              ← .env 設定讀取
├── logger/              ← Logger interface + logrus
├── db/                  ← 資料庫連線 + Migration 執行
├── migrations/          ← SQL 版本化 migration（golang-migrate）
├── seeder/              ← 初始資料填充
├── model/               ← 資料結構（Creature、Player）
├── repository/          ← 資料庫操作抽象（BaseRepository + Filters）
├── router/              ← 路由定義
├── handler/creature/    ← 生物 API 處理 + 測試
├── handler/player/      ← 玩家 API 處理 + 測試
├── service/creature/    ← 生物商業邏輯 + 測試
├── service/player/      ← 玩家商業邏輯 + 測試
└── flow/catch/          ← 捕獲生物流程（跨 service Use Case）
```

> 完整架構說明 → `.claude/docs/STRUCTURE.md`

## 開發原則

1. 新增 API 前先在 `router/router.go` 登記路由
2. 新增資料表用 SQL migration 檔案，不要直接改 AutoMigrate
3. 不在程式碼裡寫死生物資料，從圖鑑讀取參數
4. 敏感資訊放 `.env`，不進 git
5. 在方法使用一句話註解說明
