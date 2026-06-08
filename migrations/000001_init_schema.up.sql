-- ============================================================
-- Migration 000001 UP：建立初始資料表
-- ============================================================

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
