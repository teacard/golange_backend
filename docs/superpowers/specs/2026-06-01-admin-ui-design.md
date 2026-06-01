# 草原精靈獸後台 — UI 設計規範

> 狀態：已確認  
> 日期：2026-06-01  
> 設計工具：Pencil（pencil-new.pen）

---

## 概述

後台管理介面採用深色系（Dark Mode）+ 銀灰強調色，供開發者與營運人員管理遊戲資料。
架構為模組化設計，各功能單元共用同一套視覺語言，可逐步新增模組。

---

## 規劃模組

| 優先順序 | 模組 | 說明 |
|---------|------|------|
| P0 | 後台人員管理 | 帳號、角色、權限 |
| P0 | 玩家管理 | 玩家資料、物品欄 |
| P1 | 精靈管理 | 生物圖鑑維護 |
| P1 | 活動管理 | 遊戲活動設定 |
| P2 | 後台日誌 | 操作紀錄查詢 |
| P2 | 儀表板 | 統計數據概覽 |

---

## 版面結構

```
┌─────────────────────────────────────────┐
│  60px  │  Header (56px)                 │
│        ├────────────────────────────────│
│ 側欄   │  Body (padding: 24px)          │
│ Icons  │                                │
│        │  Content Area (fill)           │
│        │                                │
│────────│                                │
│ Avatar │                                │
└─────────────────────────────────────────┘
```

- **側欄**：60px 固定寬，僅顯示 icon，active 狀態背景 `#222736`
- **Header**：高 56px，左側頁面標題，右側操作區
- **Body**：padding 24px，內容區自由延伸

---

## 色彩系統

### 背景層次

| Token | Hex | 用途 |
|-------|-----|------|
| `bg-base` | `#0C0E14` | 頁面最底層背景 |
| `bg-sidebar` | `#13161E` | 側欄背景 |
| `bg-surface` | `#1A1E2A` | 卡片、面板、表格底色 |
| `bg-elevated` | `#222736` | 側欄 active、次層容器 |
| `bg-input` | `#1E2330` | 輸入框底色 |
| `bg-row-alt` | `#1D2130` | 表格交替列底色 |

### 文字

| Token | Hex | 用途 |
|-------|-----|------|
| `text-primary` | `#E8EAF0` | 主要文字、標題、數值 |
| `text-secondary` | `#9CA3AF` | 次要文字、帳號、說明 |
| `text-muted` | `#5C6478` | placeholder、欄位標籤、輔助 |

### 強調色

| Token | Hex | 用途 |
|-------|-----|------|
| `accent` | `#9CA3AF` | 側欄選中 icon、次要按鈕邊框 |
| `accent-hover` | `#C8CBD3` | 次要按鈕文字、hover 狀態 |

### 邊框

| Token | Hex | 用途 |
|-------|-----|------|
| `border` | `#262C3E` | 卡片、分隔線、表格 |
| `border-input` | `#6B7280` | 輸入框、搜尋框、下拉選單 |

### 狀態色

| 狀態 | 文字色 | 背景色 |
|------|--------|--------|
| 啟用 / Success | `#34D399` | `#052E16` |
| 警告 / Warning | `#FBBF24` | `#1C1400` |
| 錯誤 / Danger | `#F87171` | `#1C0000` |
| 資訊 / Info | `#60A5FA` | `#0A1628` |
| 停用 / Inactive | `#9CA3AF` | `#2A2A2A` |

---

## 字型規範

字型統一使用 **Inter**（Google Fonts）。

| 用途 | 大小 | 粗細 | 顏色 |
|------|------|------|------|
| 頁面標題 | 15px | 600 | `text-primary` |
| 區塊標題 | 14px | 600 | `text-primary` |
| 統計數字大標 | 24px | 700 | `text-primary` |
| 表格主要內容 | 13px | 400 | `text-primary` |
| 表格次要內容 | 13px | 400 | `text-secondary` |
| 欄位標籤 / 說明 | 12px | 500 | `text-muted` |

---

## 元件規範

### 按鈕

| 類型 | 底色 | 邊框 | 文字色 | Radius |
|------|------|------|--------|--------|
| 主要（新增） | `#374151` | — | `#E8EAF0` | 6px |
| 次要（匯出） | transparent | `#9CA3AF` | `#C8CBD3` | 6px |
| 危險（刪除） | transparent | `#F87171` | `#F87171` | 6px |

### 狀態徽章

Pill 形狀（radius: 20px），padding: `4px 10px`。
文字格式：`● 狀態文字`，12px。
顏色對照參考「狀態色」表格。

### 輸入框 / 搜尋框 / 下拉選單

- 底色：`#1E2330`
- 邊框：`#6B7280`，1px
- Radius：6px
- 高度：36px
- Placeholder 文字色：`text-muted`（`#5C6478`）
- 輸入中文字色：`text-primary`（`#E8EAF0`）

### 統計卡片

- 底色：`#13161E` 或 `#1A1E2A`
- 邊框：`border`（`#262C3E`），1px
- Radius：8px
- 高度：80–84px
- 數值字體：24px / 700
- 標籤字體：12px / `text-secondary`

### 資料表格

- 表格容器：`#1A1E2A`，radius 8px，border `#262C3E`
- 表頭列：`#161A24`，高 36px，文字 12px/500/`text-muted`
- 資料列高：48px
- 奇數列底色：`#1A1E2A`
- 偶數列底色：`#1D2130`
- 列分隔線：`border`（`#262C3E`），1px bottom

---

## 側欄導航

| 狀態 | 說明 |
|------|------|
| 預設 | icon 居中，底色透明，icon 色 `text-secondary` |
| Active | 底色 `#222736`，radius 6px，icon 色 `text-primary` |

模組 icon 尺寸：20×20px，使用方形佔位（正式開發替換為 SVG icon）。
底部常駐使用者頭像：32px 圓形，`#374151` 底色。

---

## 設計檔案

- Pencil 設計稿：`pencil-new.pen`
  - `Admin Preview`（LvMqe）：後台人員管理頁面完整預覽
  - `Style Guide`（PEgAH）：Typography / Buttons / Form / Cards 規範展示板
