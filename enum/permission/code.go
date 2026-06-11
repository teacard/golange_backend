package permission

// Code 權限識別碼型別
type Code string

const (
	/* 後台人員-檢視 */
	AdminView Code = "ADMIN_VIEW"
	/* 後台人員-編輯 */
	AdminEdit Code = "ADMIN_EDIT"
	/* 後台人員-刪除 */
	AdminDelete Code = "ADMIN_DELETE"

	/* 遊戲會員-檢視 */
	PlayerView Code = "PLAYER_VIEW"
	/* 遊戲會員-編輯 */
	PlayerEdit Code = "PLAYER_EDIT"
	/* 遊戲會員-刪除 */
	PlayerDelete Code = "PLAYER_DELETE"

	/* 公告-檢視 */
	AnnouncementView Code = "ANNOUNCEMENT_VIEW"
	/* 公告-編輯 */
	AnnouncementEdit Code = "ANNOUNCEMENT_EDIT"
	/* 公告-刪除 */
	AnnouncementDelete Code = "ANNOUNCEMENT_DELETE"
)

// LocaleKey 回傳此權限碼對應的 locale key（對應 locales/*/permission.json）
func (c Code) LocaleKey() string {
	return string(c)
}

// Cases 回傳所有已定義的權限碼
func Cases() []Code {
	return []Code{
		AdminView, AdminEdit, AdminDelete,
		PlayerView, PlayerEdit, PlayerDelete,
		AnnouncementView, AnnouncementEdit, AnnouncementDelete,
	}
}
