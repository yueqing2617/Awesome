package manager

import (
	"Awesome/app/http/helper"
	"Awesome/app/http/requests"
	"Awesome/app/models"
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"strconv"
)

type AdminRole struct {
	//Dependent services
	Model models.AdminRole
}

func NewAdminRole() *AdminRole {
	return &AdminRole{
		//Inject services
	}
}

// Index GET /dummy
func (r *AdminRole) Index(ctx http.Context) {
	page, _ := strconv.Atoi(ctx.Request().Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Request().Query("limit", "10"))
	name := ctx.Request().Query("name", "")
	order := ctx.Request().Query("order", "id desc")
	// 查询
	query := facades.Orm.Query().Model(r.Model)
	if name != "" {
		query = query.Where("name like ?", "%"+name+"%").OrWhere("display_name like ?", "%"+name+"%")
	}
	query.Order(order)
	// 分页
	type result struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Remark      string `json:"remark"`
	}
	var data []result
	var total int64
	if err := query.Paginate(page, limit, &data, &total); err != nil {
		helper.RestfulError(ctx, "查询失败，原因："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"items": data,
		"total": total,
	})
}

// Create POST /dummy
func (r *AdminRole) Create(ctx http.Context) {
	var params requests.AdminAdminRoleRequest
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "表单参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "表单参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	data := &models.AdminRole{
		Name:          params.Name,
		DisplayName:   params.DisplayName,
		Remark:        params.Remark,
		PermissionStr: helper.ArrayToString(params.Permissions),
	}
	if err := facades.Orm.Query().Model(&r.Model).Create(&data); err != nil {
		helper.RestfulError(ctx, "创建失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "创建成功", nil)
}

// Show GET /dummy/{id}
func (r *AdminRole) Show(ctx http.Context) {
	id := ctx.Request().Input("id")
	if id == "" {
		helper.RestfulError(ctx, "id:参数错误")
		return
	}
	var role models.AdminRole
	if err := facades.Orm.Query().Model(&r.Model).Where("id = ?", id).First(&role); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	if role.ID == 0 {
		helper.RestfulError(ctx, "查询失败：没有找到数据")
		return
	}
	_ = role.Scan()
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"item": http.Json{
			"name":         role.Name,
			"display_name": role.DisplayName,
			"remark":       role.Remark,
			"permissions":  role.Permissions,
		},
	})
}

// Edit GET /dummy/{id}/edit
func (r *AdminRole) Edit(ctx http.Context) {
	var params requests.AdminAdminRoleRequest
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "表单参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "表单参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	id := ctx.Request().Input("id")
	if id == "" {
		helper.RestfulError(ctx, "id:参数错误")
		return
	}
	var role models.AdminRole
	if err := facades.Orm.Query().Model(r.Model).Where("id = ?", id).First(&role); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	key := fmt.Sprintf("admin_permission_%s", role.Name)
	role.Name = params.Name
	role.DisplayName = params.DisplayName
	role.Remark = params.Remark
	role.PermissionStr = helper.ArrayToString(params.Permissions)
	res, err := facades.Orm.Query().Model(&r.Model).Where("id = ?", id).Updates(&role)
	if err != nil {
		helper.RestfulError(ctx, "更新失败："+err.Error())
		return
	}
	if res.RowsAffected == 0 {
		helper.RestfulError(ctx, "更新失败：未找到要更新的记录")
		return
	}
	// 删除缓存
	facades.Cache.Forget(key)
	helper.RestfulSuccess(ctx, "更新成功, 影响行数："+strconv.FormatInt(res.RowsAffected, 10), nil)
}

// Delete DELETE /dummy
// @param id []int
func (r *AdminRole) Delete(ctx http.Context) {
	ids := ctx.Request().QueryArray("id")
	if len(ids) == 0 {
		helper.RestfulError(ctx, "id:至少选择一条记录")
		return
	}
	var roles []models.AdminRole
	if err := facades.Orm.Query().Model(r.Model).Where("id in (?)", ids).Find(&roles); err != nil {
		helper.RestfulError(ctx, "删除失败："+err.Error())
		return
	}
	if len(roles) == 0 {
		helper.RestfulError(ctx, "删除失败：未找到要删除的记录")
		return
	}
	// 判断是否有超级管理员角色
	for _, role := range roles {
		if role.Name == "super_admin" {
			helper.RestfulError(ctx, "删除失败：不能删除超级管理员角色")
			return
		}
	}
	token := ctx.Request().Header("token", "")
	payload, err := facades.Auth.Parse(ctx, token)
	if err != nil {
		helper.RestfulError(ctx, "删除失败："+err.Error())
		return
	}
	if payload == nil {
		helper.RestfulError(ctx, "删除失败：token错误")
		return
	}
	var admin models.Admin
	if err := facades.Orm.Query().Model(admin).Where("id = ?", payload.Key).With("Role").FirstOrFail(&admin); err != nil {
		helper.RestfulError(ctx, "删除失败："+err.Error())
		return
	}
	if admin.ID == 0 {
		helper.RestfulError(ctx, "删除失败：管理员角色错误")
		return
	}
	if admin.Role.Name != "super_admin" {
		helper.RestfulError(ctx, "删除失败：只有超级管理员才能删除角色")
		return
	}
	res, err := facades.Orm.Query().Delete(&r.Model, ids)
	if err != nil {
		helper.RestfulError(ctx, "删除失败："+err.Error())
		return
	}
	// 删除缓存
	for _, role := range roles {
		key := fmt.Sprintf("admin_permission_%s", role.Name)
		facades.Cache.Forget(key)
	}

	// 成功删除了 res.RowsAffected 条记录
	helper.RestfulSuccess(ctx, "成功删除了 "+strconv.Itoa(int(res.RowsAffected))+" 条记录", nil)
}
