package seeder_test

import (
	"testing"

	"game-backend/locale"
	"game-backend/model"
	"game-backend/seeder"
	"game-backend/test/shared"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	locale.Init()
	m.Run()
}

// TestSeedPermissions_CreatesAll 驗證執行後資料庫存在 9 筆正確的 permission
func TestSeedPermissions_CreatesAll(t *testing.T) {
	// GIVEN：SQLite in-memory DB，schema 已建立
	db := shared.SetupDB(t, &model.Permission{})

	// WHEN：執行 SeedPermissions
	seeder.SeedPermissions(db, &shared.NopLogger{})

	// THEN：共 9 筆資料
	var count int64
	db.Table("permissions").Count(&count)
	assert.Equal(t, int64(9), count)
}

// TestSeedPermissions_Idempotent 驗證重複執行兩次不產生重複資料列
func TestSeedPermissions_Idempotent(t *testing.T) {
	// GIVEN：SQLite in-memory DB，第一次執行已完成
	db := shared.SetupDB(t, &model.Permission{})
	seeder.SeedPermissions(db, &shared.NopLogger{})

	// WHEN：再次執行 SeedPermissions
	seeder.SeedPermissions(db, &shared.NopLogger{})

	// THEN：仍只有 9 筆，不重複寫入
	var count int64
	db.Table("permissions").Count(&count)
	assert.Equal(t, int64(9), count)
}
