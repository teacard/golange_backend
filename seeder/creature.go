package seeder

import (
	"game-backend/db"
	"game-backend/logger"
	"game-backend/model"

	"gorm.io/datatypes"
)

// SeedCreatures 寫入圖鑑內的初始生物資料
// 使用 FirstOrCreate：若已存在則略過，避免重複執行時出錯
func SeedCreatures(log logger.Logger) {
	log.Info("Seeder：生物資料...")

	seedGrassSpirit(log)

	// TODO: 新增更多生物時在此追加 seed 函式
}

// seedGrassSpirit 草原精靈獸（第一隻，風格基準）
// 圖鑑來源：claude/creatures/grassspirit.md
func seedGrassSpirit(log logger.Logger) {
	// Step 1：建立或取得生物主資料
	// FirstOrCreate：
	//   → 先用 Where 條件查詢，找到就回傳現有資料（RowsAffected = 0）
	//   → 找不到才 INSERT（RowsAffected = 1）
	//   → 無論哪種情況，creature.ID 執行後都會有值
	creature := model.Creature{
		CreatureID:  "creature_grassspirit",
		Rarity:      "uncommon",
		Personality: "悠閒型",
		Habitat:     "草原",
		Colors: datatypes.JSON(`{
			"primary":   "#e8d8b0",
			"secondary": "#b8a070"
		}`),
		SpecialParts: datatypes.JSON(`[
			{"code": "head_leaf", "color": "#8aba60"},
			{"code": "round_ears"}
		]`),
		Animation: datatypes.JSON(`{
			"idle":  {"breath_speed": 1.2, "breath_amp": 3,  "tail_speed": 0.9, "tail_amp": 0.12, "blink_interval": 0.8},
			"walk":  {"cycle": 3.0, "body_bounce": 5, "body_sway": 0.04, "leg_amp": 0.35, "tail_amp": 0.2},
			"catch": {"shake_freq": 18, "shake_amp": 8, "bounce_amp": 10, "duration": 1.5, "blush_alpha": 0.6}
		}`),
	}

	result := db.DB.Where(model.Creature{CreatureID: "creature_grassspirit"}).FirstOrCreate(&creature)
	if result.Error != nil {
		log.Infof("Seeder：草原精靈獸寫入失敗 - %v", result.Error)
		return
	}
	if result.RowsAffected > 0 {
		log.Info("Seeder：草原精靈獸已寫入")
	} else {
		log.Info("Seeder：草原精靈獸已存在，略過")
	}

	// Step 2：寫入語系資料
	// FK 改用 creature.ID（uint，DB primary key）
	// 必須在 Step 1 之後才能取得 creature.ID
	locales := []model.CreatureLocale{
		{
			CreatureID:  creature.ID, // DB primary key（uint），不是 creature_id 字串
			LangCode:    "zh-TW",
			Name:        "草原精靈獸",
			Description: "棲息於廣闊草原的溫馴精靈獸，頭頂插著一片能感應天氣的嫩葉。",
		},
		{
			CreatureID:  creature.ID, // 同上
			LangCode:    "en",
			Name:        "Grass Spirit",
			Description: "A gentle spirit beast of the vast grasslands, with a leaf atop its head said to sense the weather.",
		},
	}

	for _, locale := range locales {
		db.DB.Where(model.CreatureLocale{
			CreatureID: locale.CreatureID,
			LangCode:   locale.LangCode,
		}).FirstOrCreate(&locale)
	}
}
