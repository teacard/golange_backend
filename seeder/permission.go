package seeder

import (
	"game-backend/enum/permission"
	"game-backend/locale"
	"game-backend/logger"
	"game-backend/model"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

/*
	SeedPermissions 寫入系統預定義權限初始資料
	使用 FirstOrCreate：以 code 為條件，已存在則略過，冪等安全
*/
func SeedPermissions(gormDB *gorm.DB, log logger.Logger) {
	log.Info("Seeder：權限資料...")

	// seeder 寫入固定用 zh-TW，不需要 gin context
	l := locale.Localizer("zh-TW")

	for _, code := range permission.Cases() {
		name, _ := l.Localize(&i18n.LocalizeConfig{MessageID: code.LocaleKey()})
		perm := model.Permission{Name: name, Code: code}
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
