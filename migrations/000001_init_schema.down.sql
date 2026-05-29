-- ============================================================
-- Migration 000001 DOWN：回滾初始資料表（刪除全部）
-- ============================================================

DROP TABLE IF EXISTS inventories;
DROP TABLE IF EXISTS players;
DROP TABLE IF EXISTS creature_locales;
DROP TABLE IF EXISTS creatures;
