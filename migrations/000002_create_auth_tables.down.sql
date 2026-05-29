-- ============================================================
-- Migration 000002 DOWN：回滾 Auth 相關資料表
-- 依照 FK 依賴順序反向刪除
-- ============================================================

DROP TABLE IF EXISTS role_ables;
DROP TABLE IF EXISTS model_roles;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS admins;
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS permissions;
