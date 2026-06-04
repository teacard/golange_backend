---
name: git-commit
description: Use when making any git commit in this project — requires creating a branch first, never commit directly to develop or main.
---

# Git Commit 規範

## 核心原則

**永遠不直接 commit 到 `develop` 或 `main`，必須先建立分支。**

## 流程

```
1. 確認目前分支
2. 建立新分支（若還在 develop/main）
3. 進行變更
4. Stage 指定檔案
5. 寫 commit message
```

### 1. 確認目前分支

```bash
git branch
```

若在 `develop` 或 `main`，必須先切新分支才能繼續。

### 2. 建立分支

```bash
git checkout -b <type>/<description>
```

| type | 用途 |
|------|------|
| `feat` | 新功能 |
| `fix` | 修復 bug |
| `docs` | 文件 |
| `refactor` | 重構 |
| `test` | 測試 |
| `chore` | 維護、設定 |

範例：`feat/player-inventory-api`、`docs/admin-ui-design`

### 3. Stage 指定檔案

```bash
git add <file1> <file2>
```

不使用 `git add .` 或 `git add -A`，避免誤加 `.env` 或其他敏感檔案。

### 4. Commit Message 格式

```
<type>: <動詞開頭的中文描述>

<選填：補充說明>

Co-Authored-By: Claude Sonnet 4.6 <noreply@anthropic.com>
```

範例：
```
feat: 新增玩家物品欄 API

支援 GET /api/players/:id/inventory，回傳分頁結果。

Co-Authored-By: Claude Sonnet 4.6 <noreply@anthropic.com>
```

## 紅旗 — 立即停止

- 目前在 `develop` 或 `main` 就想要 commit → 先建分支
- 準備用 `git add .` → 改為指定檔案
- commit message 沒有 type prefix → 補上
