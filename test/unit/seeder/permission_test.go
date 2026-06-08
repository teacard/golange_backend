package seeder_test

import (
	"testing"

	"game-backend/seeder"
	"game-backend/test/shared"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// TestSeedPermissions_CallsFirstOrCreateForEachPermission 驗證每個 permission 都會觸發一次 FirstOrCreate
func TestSeedPermissions_CallsFirstOrCreateForEachPermission(t *testing.T) {
	// GIVEN：空的 permissions 資料表（SELECT 回傳空列）
	gormDB, mock := shared.NewMockGormDB(t)
	for i := 1; i <= 9; i++ {
		mock.ExpectQuery(`SELECT`).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "code"}))
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "permissions"`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(i))
		mock.ExpectCommit()
	}

	// WHEN：執行 SeedPermissions
	seeder.SeedPermissions(gormDB, &shared.NopLogger{})

	// THEN：9 筆 SELECT + INSERT 都被呼叫
	assert.NoError(t, mock.ExpectationsWereMet())
}

// TestSeedPermissions_IdempotentWhenAlreadyExists 驗證資料已存在時不重複寫入
func TestSeedPermissions_IdempotentWhenAlreadyExists(t *testing.T) {
	// GIVEN：permissions 資料表已有所有 9 筆資料（SELECT 回傳現有列）
	gormDB, mock := shared.NewMockGormDB(t)
	cols := []string{"id", "created_at", "updated_at", "name", "code"}
	for i := 1; i <= 9; i++ {
		mock.ExpectQuery(`SELECT`).
			WillReturnRows(sqlmock.NewRows(cols).AddRow(i, nil, nil, "name", "code"))
	}

	// WHEN：執行 SeedPermissions
	seeder.SeedPermissions(gormDB, &shared.NopLogger{})

	// THEN：只有 SELECT，沒有任何 INSERT（資料已存在，FirstOrCreate 略過）
	assert.NoError(t, mock.ExpectationsWereMet())
}
