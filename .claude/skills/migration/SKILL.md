---
name: migration
description: >
  草原精靈獸遊戲 — PostgreSQL Migration SQL 撰寫規範。
  每次新增或修改資料表時必須使用此規範。
  包含：資料表結構標準、欄位順序、主鍵規範、註解格式、命名規則。
---

# Migration SQL 撰寫規範

## 檔案命名

```
migrations/
├── {版本號}_{說明}.up.sql    ← 建立 / 修改
└── {版本號}_{說明}.down.sql  ← 回滾（必須完整還原 up 的動作）
```

版本號從 `000001` 開始，每次遞增，永遠補足六位數。

```
000001_init_schema.up.sql
000002_add_creature_level.up.sql
000003_add_player_stamina.up.sql
```

---

## 資料表結構標準

### 欄位順序（固定）

```
1. id               ← 主鍵，永遠第一欄
2. [業務欄位...]    ← 該表核心資料
3. created_at       ← 建立時間，永遠倒數第三
4. updated_at       ← 更新時間，永遠倒數第二
5. deleted_at       ← 軟刪除時間，永遠最後一欄
```

### 主鍵規範

- 型別一律使用 `SERIAL`（對應 INT，4 bytes，約 21 億上限）
- 遊戲規模不需要 BIGSERIAL（8 bytes）
- 名稱一律為 `id`

```sql
id SERIAL PRIMARY KEY,
```

### 時間戳記規範

- 型別一律使用 `TIMESTAMPTZ`（含時區）
- `created_at / updated_at` 所有表都需要，放在最末端
- `deleted_at` 只放在**主表**，子表不需要

```sql
-- 主表（可獨立軟刪除）
created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
deleted_at  TIMESTAMPTZ           -- NULL = 未刪除

-- 子表（透過 JOIN 父表過濾，不需要自己的 deleted_at）
created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
```

### 主表 vs 子表的判斷

| 情況 | 需要 `deleted_at`？ | 理由 |
|------|---------------------|------|
| 可以被獨立停用 / 下架的表 | ✅ 是 | 例：`creatures`、`players` |
| 內容隨父表存在的附屬表 | ❌ 否 | 例：`creature_locales`（生物不存在，語系資料也沒意義） |
| 業務行為決定（如背包） | 視需求 | 若「放生」需保留紀錄 → 要；若直接刪除 → 不要 |

子表的可見性由查詢時 JOIN 父表的 `WHERE deleted_at IS NULL` 控制，不需要在子表重複紀錄。

---

## FK（外鍵）設計原則

**全系統統一：所有外鍵一律指向 DB primary key（`id INTEGER`）**

```sql
-- ✅ 正確：FK 指向 creatures.id（DB primary key）
creature_id INTEGER NOT NULL REFERENCES creatures(id)

-- ❌ 錯誤：FK 指向 creatures.creature_id（business id，varchar）
creature_id VARCHAR(100) NOT NULL REFERENCES creatures(creature_id)
```

### 為什麼這樣設計？

| 比較 | 指向 DB PK（int） | 指向 business id（varchar） |
|------|-----------------|--------------------------|
| 查詢效能 | 整數比較，最快 | 字串比較，較慢 |
| 儲存空間 | 4 bytes | 最多 100 bytes |
| 一致性 | 全系統統一 | 混用導致 schema inconsistency |

### Business ID 的用途

`creature_id VARCHAR`、`player_id VARCHAR` 這類業務識別碼仍然保留在各自的主表，用途是：
- **API 對外**：URL 路徑、JSON response 使用業務識別碼，不暴露資料庫 PK
- **Seeder**：用業務識別碼做 `FirstOrCreate` 條件，語意清楚
- **不做 FK**：關聯一律用整數 PK，業務識別碼只供查詢條件

---

## 註解規範

**每張資料表**必須有兩種註解：

### 1. 表層級註解（說明這張表的用途）

```sql
COMMENT ON TABLE creatures IS '生物主表：遊戲中所有可蒐集生物的基本資料';
```

### 2. 欄位層級註解（說明欄位用途與允許值）

```sql
COMMENT ON COLUMN creatures.rarity IS '稀有度：common / uncommon / rare / legendary';
COMMENT ON COLUMN creatures.colors IS 'JSONB：{"primary":"#hex","secondary":"#hex"}';
```

- 如果欄位名稱語意已完全清楚（如 `id`、`created_at`），可省略欄位註解
- JSONB 欄位**必須**在註解裡說明預期格式

---

## 完整資料表範本

```sql
-- ============================================================
-- {表名}：{一句話說明用途}
-- ============================================================
CREATE TABLE IF NOT EXISTS {table_name} (
    id          SERIAL      PRIMARY KEY,

    -- 業務欄位
    some_id     VARCHAR(100) NOT NULL UNIQUE,
    some_field  VARCHAR(200) NOT NULL,
    some_json   JSONB        NOT NULL DEFAULT '{}',

    -- 時間戳記（固定放最後）
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);

-- 索引
CREATE INDEX IF NOT EXISTS idx_{table_name}_deleted_at ON {table_name}(deleted_at);

-- 表註解
COMMENT ON TABLE {table_name} IS '{表的用途說明}';

-- 欄位註解（有疑義或 JSONB 欄位才需要）
COMMENT ON COLUMN {table_name}.some_json IS 'JSONB：{預期格式}';
```

---

## 索引命名規則

```
idx_{資料表名}_{欄位名}
idx_{資料表名}_{欄位1}_{欄位2}   ← 複合索引
```

範例：
```sql
CREATE INDEX IF NOT EXISTS idx_creatures_deleted_at    ON creatures(deleted_at);
CREATE INDEX IF NOT EXISTS idx_creatures_creature_id   ON creatures(creature_id);
CREATE INDEX IF NOT EXISTS idx_creature_locales_lookup ON creature_locales(creature_id, lang_code);
```

---

## Down 檔案規範

Down 必須完整還原 Up 的所有動作，順序相反（先刪依賴表，再刪被依賴表）。

```sql
-- ============================================================
-- Migration {版本號} DOWN：還原 {說明}
-- ============================================================
DROP TABLE IF EXISTS {依賴表};   -- 先刪有外鍵的
DROP TABLE IF EXISTS {主表};     -- 後刪被參考的
```

---

## 新增資料表 Checklist

撰寫新 migration 前逐項確認：

- [ ] 檔案版本號接續上一個
- [ ] 每張表第一欄是 `id SERIAL PRIMARY KEY`
- [ ] 業務欄位在中間
- [ ] `created_at / updated_at` 排在最後
- [ ] 主表才加 `deleted_at`；子表不加
- [ ] 有 `deleted_at` 的表才建 `idx_{table}_deleted_at` 索引
- [ ] `COMMENT ON TABLE` 有填寫
- [ ] JSONB 欄位有 `COMMENT ON COLUMN` 說明格式
- [ ] FK 欄位型別是 `INTEGER`，指向對方的 `id`（DB primary key），不是 business id
- [ ] Down 檔案可完整回滾
