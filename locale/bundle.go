package locale

import (
	"embed"
	"encoding/json"
	"io/fs"
	pathpkg "path"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed locales
var localeFS embed.FS

var Bundle *i18n.Bundle

func Init() {
	Bundle = i18n.NewBundle(language.TraditionalChinese)
	Bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// 遞迴掃描 locales/ 下所有子目錄的 JSON 檔
	// 語言 tag 由目錄名稱決定（zh-TW/ → zh-TW，en/ → en）
	if err := fs.WalkDir(localeFS, "locales", func(filePath string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		data, err := localeFS.ReadFile(filePath)
		if err != nil {
			panic("locale: 無法讀取語系檔 " + filePath + "：" + err.Error())
		}
		// go-i18n 從「倒數第二個點分段」取語言 tag
		// 正確格式：category.lang.json（例如 api_code.zh-TW.json）
		// embed.FS 路徑固定使用正斜線，用 path 套件而非 filepath
		lang := pathpkg.Base(pathpkg.Dir(filePath))
		category := strings.TrimSuffix(pathpkg.Base(filePath), ".json")
		parsePath := category + "." + lang + ".json"
		if _, err := Bundle.ParseMessageFileBytes(data, parsePath); err != nil {
			panic("locale: 解析語系檔 " + filePath + " 失敗：" + err.Error())
		}
		return nil
	}); err != nil {
		panic("locale: 無法掃描語系目錄：" + err.Error())
	}
}

// Localizer 建立指定語系的 Localizer
func Localizer(langs ...string) *i18n.Localizer {
	return i18n.NewLocalizer(Bundle, langs...)
}
