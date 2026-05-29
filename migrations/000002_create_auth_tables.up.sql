-- ============================================================
-- Migration 000002 UP：建立 Auth 相關資料表
-- 包含：permissions, roles, role_permissions, admins, users, model_roles
-- ============================================================


-- ── permissions ──────────────────────────────────────────────
-- 權限清單：以 code 欄位作為 enum 控制依據
CREATE TABLE IF NOT EXISTS permissions (
    id         SERIAL       PRIMARY KEY,
    name       VARCHAR(100) NOT NULL,
    code       VARCHAR(100) NOT NULL UNIQUE,

    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_permissions_code ON permissions(code);

COMMENT ON TABLE  permissions      IS '權限清單：定義系統中所有可控制的操作';
COMMENT ON COLUMN permissions.name IS '顯示名稱，例如：上傳圖鑑';
COMMENT ON COLUMN permissions.code IS '程式識別碼（enum 用），例如：UPLOAD_CREATURE';


-- ── roles ────────────────────────────────────────────────────
-- 角色清單：同時供 admins 與 users 使用
CREATE TABLE IF NOT EXISTS roles (
    id         SERIAL       PRIMARY KEY,
    name       VARCHAR(100) NOT NULL UNIQUE,

    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE  roles      IS '角色清單：admins 與 users 共用';
COMMENT ON COLUMN roles.name IS '角色名稱，例如：super_admin / editor';


-- ── role_permissions ─────────────────────────────────────────
-- 角色權限對應表（pivot）：一個 role 對應多個 permission
CREATE TABLE IF NOT EXISTS role_permissions (
    role_id       INTEGER NOT NULL REFERENCES roles(id),
    permission_id INTEGER NOT NULL REFERENCES permissions(id),

    UNIQUE (role_id, permission_id)
);

CREATE INDEX IF NOT EXISTS idx_role_permissions_role_id ON role_permissions(role_id);

COMMENT ON TABLE role_permissions IS '角色 ↔ 權限對應（pivot）';


-- ── admins ───────────────────────────────────────────────────
-- 後台管理員帳號表
CREATE TABLE IF NOT EXISTS admins (
    id            SERIAL       PRIMARY KEY,
    email         VARCHAR(255) NOT NULL UNIQUE,
    password      VARCHAR(255) NOT NULL,
    name          VARCHAR(200) NOT NULL,
    status        VARCHAR(50)  NOT NULL DEFAULT 'active',
    last_login_at TIMESTAMPTZ,

    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_admins_deleted_at ON admins(deleted_at);
CREATE INDEX IF NOT EXISTS idx_admins_email      ON admins(email);

COMMENT ON TABLE  admins             IS '後台管理員帳號表';
COMMENT ON COLUMN admins.status      IS '帳號狀態：active / suspended';
COMMENT ON COLUMN admins.last_login_at IS '最後登入時間';


-- ── users ────────────────────────────────────────────────────
-- 前台玩家帳號表
CREATE TABLE IF NOT EXISTS users (
    id            SERIAL       PRIMARY KEY,
    user_id       VARCHAR(26)  NOT NULL UNIQUE,  -- ULID 固定 26 字元
    email         VARCHAR(255) NOT NULL UNIQUE,
    password      VARCHAR(255) NOT NULL,
    nickname      VARCHAR(200) NOT NULL,
    status        VARCHAR(50)  NOT NULL DEFAULT 'active',
    last_login_at TIMESTAMPTZ,

    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users(deleted_at);
CREATE INDEX IF NOT EXISTS idx_users_user_id    ON users(user_id);
CREATE INDEX IF NOT EXISTS idx_users_email      ON users(email);

COMMENT ON TABLE  users              IS '前台玩家帳號表';
COMMENT ON COLUMN users.user_id      IS 'ULID，對外識別碼，固定 26 字元';
COMMENT ON COLUMN users.status       IS '帳號狀態：active / suspended';
COMMENT ON COLUMN users.last_login_at IS '最後登入時間';


-- ── model_roles ──────────────────────────────────────────────
-- 多型角色對應表（polymorphic pivot）：admins 與 users 共用
-- model 欄位區分來源資料表（admin / user）
CREATE TABLE IF NOT EXISTS model_roles (
    model    VARCHAR(50) NOT NULL,
    model_id INTEGER     NOT NULL,
    role_id  INTEGER     NOT NULL REFERENCES roles(id),

    UNIQUE (model, model_id, role_id)
);

CREATE INDEX IF NOT EXISTS idx_model_roles_lookup ON model_roles(model, model_id);

COMMENT ON TABLE  model_roles          IS '多型角色對應（polymorphic pivot）：admins 與 users 共用';
COMMENT ON COLUMN model_roles.model    IS '來源 model 名稱：admin / user';
COMMENT ON COLUMN model_roles.model_id IS '對應各自資料表的 id';
