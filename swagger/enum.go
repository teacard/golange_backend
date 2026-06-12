package swagger

import enumresponse "game-backend/dto/enum/response"

// @Summary      業務錯誤碼對照表
// @Description  列出所有 ApiCode 的值與對應中文說明，供前端做錯誤碼渲染用
// @Tags         Enum與對照表
// @Produce      json
// @Success      200  {array}  enumresponse.ApiCodeItem
// @Router       /admin-api/enums/api-code [get]
func ApiCodeEnumDoc(_ enumresponse.ApiCodeItem) {}
