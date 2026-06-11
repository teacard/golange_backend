package api_code

// ApiCode 業務錯誤識別碼型別
type ApiCode string

const (
	// 角色名稱已被使用
	RoleNameTaken ApiCode = "ROLE_NAME_TAKEN"
	// 角色仍有帳號使用中
	RoleInUse ApiCode = "ROLE_IN_USE"
)

// AllApiCodes 所有業務錯誤碼，新增 const 時同步在此補上
var AllApiCodes = []ApiCode{
	RoleNameTaken,
	RoleInUse,
}

// LocaleKey 回傳此錯誤碼對應的 locale key（對應 locales/*/api_code.json）
func (c ApiCode) LocaleKey() string {
	return string(c)
}

