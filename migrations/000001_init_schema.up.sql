-- ============================================================
-- Migration 000001 UP：建立初始四張資料表
-- FK 設計原則：所有外鍵一律指向 DB primary key（id INTEGER）
-- ============================================================

-- ── creatures ───────────────────────────────────────────────
-- 生物主表：遊戲中所有可蒐集生物的基本屬性資料
CREATE TABLE IF NOT EXISTS creatures (
    id            SERIAL       PRIMARY KEY,

    -- 業務識別碼（供 API 與 seeder 使用，不做 FK 關聯）
    creature_id   VARCHAR(100) NOT NULL UNIQUE,

    -- 分類屬性
    rarity        VARCHAR(50)  NOT NULL,
    personality   VARCHAR(100),
    habitat       VARCHAR(100),

    -- 外觀資料（JSONB 格式）
    colors        JSONB        NOT NULL DEFAULT '{}',
    special_parts JSONB                 DEFAULT '[]',
    animation     JSONB                 DEFAULT '{}',

    -- 時間戳記
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_creatures_deleted_at  ON creatures(deleted_at);
CREATE INDEX IF NOT EXISTS idx_creatures_creature_id ON creatures(creature_id);

COMMENT ON TABLE  creatures               IS '生物主表：遊戲中所有可蒐集生物的基本屬性資料';
COMMENT ON COLUMN creatures.creature_id   IS '業務識別碼，格式：creature_{英文名}，供 API 與 seeder 使用';
COMMENT ON COLUMN creatures.rarity        IS '稀有度：common / uncommon / rare / legendary';
COMMENT ON COLUMN creatures.personality   IS '個性類型，影響動畫行為';
COMMENT ON COLUMN creatures.habitat       IS '棲息地，用於篩選與活動地點';
COMMENT ON COLUMN creatures.colors        IS 'JSONB：{"primary":"#hex","secondary":"#hex"}';
COMMENT ON COLUMN creatures.special_parts IS 'JSONB：[{"code":"head_leaf","color":"#hex"},...]';
COMMENT ON COLUMN creatures.animation     IS 'JSONB：{"idle":{...},"walk":{...},"catch":{...}}';


-- ── creature_locales ─────────────────────────────────────────
-- 生物多語系表：各語言的名稱與描述文字
-- 子表，不需要 deleted_at：透過 JOIN creatures WHERE deleted_at IS NULL 過濾
-- FK：creature_id → creatures.id（DB primary key，不是 creature_id varchar）
CREATE TABLE IF NOT EXISTS creature_locales (
    id          SERIAL      PRIMARY KEY,

    -- FK 指向 creatures.id（DB primary key）
    creature_id INTEGER     NOT NULL REFERENCES creatures(id),

    lang_code   VARCHAR(20) NOT NULL,
    name        VARCHAR(200) NOT NULL,
    description TEXT,

    -- 時間戳記（子表不需要 deleted_at）
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (creature_id, lang_code)
);

CREATE INDEX IF NOT EXISTS idx_creature_locales_lookup ON creature_locales(creature_id, lang_code);

COMMENT ON TABLE  creature_locales             IS '生物多語系表：各語言的顯示名稱與描述文字';
COMMENT ON COLUMN creature_locales.creature_id IS 'FK → creatures.id（DB primary key）';
COMMENT ON COLUMN creature_locales.lang_code   IS '語系代碼：zh-TW / en / ja ...';


-- ── players ──────────────────────────────────────────────────
-- 玩家主表：遊戲玩家的帳號資料
CREATE TABLE IF NOT EXISTS players (
    id        SERIAL       PRIMARY KEY,

    -- 業務識別碼（供 API 使用，例如 UUID）
    player_id VARCHAR(100) NOT NULL UNIQUE,
    name      VARCHAR(200) NOT NULL,

    -- 時間戳記
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_players_deleted_at ON players(deleted_at);
CREATE INDEX IF NOT EXISTS idx_players_player_id  ON players(player_id);

COMMENT ON TABLE  players           IS '玩家主表：遊戲玩家的帳號與顯示名稱';
COMMENT ON COLUMN players.player_id IS '業務識別碼（UUID），供 API 與認證使用';
COMMENT ON COLUMN players.name      IS '玩家顯示名稱';


-- ── inventories ──────────────────────────────────────────────
-- 玩家背包表：玩家已蒐集的生物清單
-- FK：全部指向 DB primary key（id），統一關聯策略
CREATE TABLE IF NOT EXISTS inventories (
    id          SERIAL      PRIMARY KEY,

    -- FK 指向 players.id（DB primary key）
    player_id   INTEGER     NOT NULL REFERENCES players(id),

    -- FK 指向 creatures.id（DB primary key）
    creature_id INTEGER     NOT NULL REFERENCES creatures(id),

    caught_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    -- 時間戳記
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ,

    UNIQUE (player_id, creature_id)
);

CREATE INDEX IF NOT EXISTS idx_inventories_deleted_at ON inventories(deleted_at);
CREATE INDEX IF NOT EXISTS idx_inventories_player_id  ON inventories(player_id);

COMMENT ON TABLE  inventories             IS '玩家背包表：記錄每位玩家已蒐集的生物與捕獲時間';
COMMENT ON COLUMN inventories.player_id   IS 'FK → players.id（DB primary key）';
COMMENT ON COLUMN inventories.creature_id IS 'FK → creatures.id（DB primary key）';
COMMENT ON COLUMN inventories.caught_at   IS '生物被捕獲的時間';
