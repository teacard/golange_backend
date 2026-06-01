---
name: git-pr
description: Use when creating a pull request to merge a feature branch back to develop — push branch first, then use gh pr create with proper body format.
---

# Git PR 規範

## 核心原則

PR 一律合併回 `develop`，不直接合併到 `main`。
Merge 後分支自動刪除（repo 已設定 `delete_branch_on_merge: true`）。

## 流程

```
1. 推送分支到遠端
2. 建立 PR（gh pr create）
3. 等待 review 後 merge
```

### 1. 推送分支

```bash
git push -u origin <branch-name>
```

### 2. 建立 PR

```bash
gh pr create --title "<type>: <標題>" --body "$(cat <<'EOF'
## Summary

- <變更重點一>
- <變更重點二>

## Test plan

- [ ] <測試項目一>
- [ ] <測試項目二>

🤖 Generated with [Claude Code](https://claude.com/claude-code)
EOF
)"
```

### PR 標題格式

與 commit message 相同：`<type>: <中文描述>`

範例：`feat: 新增玩家物品欄 API`

### Summary 區塊

- 條列主要變更，每條一行
- 說明「做了什麼」，不解釋「為什麼要這樣寫」

### Test plan 區塊

- 列出可驗證的測試項目
- 使用 `- [ ]` checkbox 格式

## 遠端設定備註

- Remote：`https://github.com/teacard/golange_backend.git`
- Base branch：`develop`
- 自動刪除分支：已啟用

## 紅旗 — 立即停止

- 分支還沒 push 就執行 `gh pr create` → 先 push
- Base branch 設成 `main` → 改為 `develop`
- PR 標題沒有 type prefix → 補上
