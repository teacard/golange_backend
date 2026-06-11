package handler

import (
	"errors"
	"game-backend/middleware"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// Localize 從 gin.Context 取得 localizer 並翻譯 messageID，失敗時回傳 messageID 本身
func Localize(ctx *gin.Context, messageID string) string {
	raw, exists := ctx.Get(middleware.LocalizerKey)
	if !exists {
		return messageID
	}
	msg, err := raw.(*i18n.Localizer).Localize(&i18n.LocalizeConfig{MessageID: messageID})
	if err != nil {
		return messageID
	}
	return msg
}

// LocalizeWithData 同 Localize，額外支援 template 插值資料
func LocalizeWithData(ctx *gin.Context, messageID string, data map[string]any) string {
	raw, exists := ctx.Get(middleware.LocalizerKey)
	if !exists {
		return messageID
	}
	msg, err := raw.(*i18n.Localizer).Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: data,
	})
	if err != nil {
		return messageID
	}
	return msg
}

// TranslateValidationError 將 validator.ValidationErrors 轉為本地化訊息
// fieldKeys 由各 request struct 的 FieldLocaleKeys() 提供，對應 struct 欄位名稱 → locale key
func TranslateValidationError(ctx *gin.Context, err error, fieldKeys map[string]string) string {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) || len(ve) == 0 {
		return err.Error()
	}
	fe := ve[0]
	fieldLocaleKey, ok := fieldKeys[fe.Field()]
	if !ok {
		return fe.Error()
	}
	return LocalizeWithData(ctx, fe.Tag(), map[string]any{
		"Field": Localize(ctx, fieldLocaleKey),
		"Param": fe.Param(),
	})
}
