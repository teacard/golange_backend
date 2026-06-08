package seeder

import (
	"game-backend/enum/permission"
	"game-backend/logger"
	"game-backend/model"

	"gorm.io/gorm"
)

/*
	SeedPermissions 寫入系統預定義權限初始資料
	使用 FirstOrCreate：以 code 為條件，已存在則略過，冪等安全
*/
func SeedPermissions(gormDB *gorm.DB, log logger.Logger) {
	log.Info("Seeder：權限資料...")

	for _, code := range permission.Cases() {
		perm := model.Permission{Name: code.Name(), Code: code}
		result := gormDB.Where(model.Permission{Code: perm.Code}).FirstOrCreate(&perm)
		if result.Error != nil {
			log.Infof("Seeder：權限 %s 寫入失敗 - %v", code, result.Error)
			continue
		}
		if result.RowsAffected > 0 {
			log.Infof("Seeder：權限 %s 已寫入", code)
		}
	}

	log.Info("Seeder：權限資料完成")
}
