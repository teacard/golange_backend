package swagger

import (
	rolerequest "game-backend/dto/role/request"
	roleresponse "game-backend/dto/role/response"
	"game-backend/exception"
)

// @Summary      角色列表
// @Tags         角色管理
// @Security     BearerAuth
// @Produce      json
// @Param        page     query  int     false  "頁碼（預設 1）"
// @Param        perPage  query  int     false  "每頁筆數（預設 20）"
// @Param        name     query  string  false  "名稱關鍵字（模糊搜尋）"
// @Success      200  {object}  roleresponse.RoleListResponse
// @Failure      422  {object}  exception.BaseErrorResponse
// @Router       /admin-api/roles [get]
func RoleList(_ roleresponse.RoleListResponse, _ exception.BaseErrorResponse) {}

// @Summary      取得單筆角色
// @Tags         角色管理
// @Security     BearerAuth
// @Produce      json
// @Param        id   path  int  true  "角色 ID"
// @Success      200  {object}  roleresponse.RoleResponse
// @Failure      404  {object}  exception.BaseErrorResponse
// @Router       /admin-api/roles/{id} [get]
func RoleGetByID(_ roleresponse.RoleResponse, _ exception.BaseErrorResponse) {}

// @Summary      建立角色
// @Tags         角色管理
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body  body  rolerequest.CreateRoleRequest  true  "角色資料"
// @Success      200  {object}  roleresponse.RoleResponse
// @Failure      422  {object}  exception.UnprocessableEntityResponse
// @Router       /admin-api/roles [post]
func RoleCreate(_ rolerequest.CreateRoleRequest, _ roleresponse.RoleResponse, _ exception.UnprocessableEntityResponse) {}

// @Summary      更新角色
// @Tags         角色管理
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id    path  int                            true  "角色 ID"
// @Param        body  body  rolerequest.UpdateRoleRequest  true  "角色資料"
// @Success      200  {object}  roleresponse.RoleResponse
// @Failure      404  {object}  exception.BaseErrorResponse
// @Failure      422  {object}  exception.UnprocessableEntityResponse
// @Router       /admin-api/roles/{id} [put]
func RoleUpdate(_ rolerequest.UpdateRoleRequest, _ roleresponse.RoleResponse, _ exception.BaseErrorResponse, _ exception.UnprocessableEntityResponse) {}

// @Summary      刪除角色
// @Tags         角色管理
// @Security     BearerAuth
// @Produce      json
// @Param        id   path  int  true  "角色 ID"
// @Success      200
// @Failure      404  {object}  exception.BaseErrorResponse
// @Failure      422  {object}  exception.UnprocessableEntityResponse
// @Router       /admin-api/roles/{id} [delete]
func RoleDelete(_ exception.BaseErrorResponse, _ exception.UnprocessableEntityResponse) {}
