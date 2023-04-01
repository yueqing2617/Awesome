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

type AdminPermission struct {
	//Dependent services
	Model models.AdminPermission
}

func NewAdminPermission() *AdminPermission {
	return &AdminPermission{
		//Inject services
	}
}

// Index GET /dummy
func (r *AdminPermission) Index(ctx http.Context) {
	opt := ctx.Request().Query("type", "tree")
	var data []models.PermissionResult
	var total int64
	var err error
	if opt == "tree" {
		data, err = (new(models.AdminPermission)).GetPermissionList(0)
		if err != nil {
			helper.RestfulError(ctx, "查询错误，原因: "+err.Error())
			return
		}
		total = int64(len(data))
	} else if opt == "list" {
		var data []models.AdminPermission
		var total int64
		if err := facades.Orm.Query().Find(&data); err != nil {
			helper.RestfulError(ctx, "查询错误，原因: "+err.Error())
			return
		}
		if err := facades.Orm.Query().Count(&total); err != nil {
			helper.RestfulError(ctx, "查询错误，原因: "+err.Error())
			return
		}
	}
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"items": data,
		"total": total,
	})

}

// Create POST /dummy
func (r *AdminPermission) Create(ctx http.Context) {
	var params requests.AdminAdminPermissionRequest
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "表单参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "表单参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	data := &models.AdminPermission{
		Name:      params.Name,
		Path:      params.Path,
		ParentId:  params.ParentId,
		Remark:    params.Remark,
		Redirect:  params.Redirect,
		Component: params.Component,
		Active:    params.Active,
		MetaJSON:  helper.PermissionMetaToJSON(&params.Meta),
	}
	fmt.Println(data.MetaJSON, params.Meta)
	if err := facades.Orm.Query().Create(&data); err != nil {
		helper.RestfulError(ctx, "创建失败，原因："+err.Error())
		return
	}
	// 如果是按钮，添加到admin_permission_policy
	if params.Meta.Type == "button" {
		policy := &models.AdminPermissionPolicy{
			Name: params.Name,
			Path: params.Path,
		}
		if err := facades.Orm.Query().Create(&policy); err != nil {
			helper.RestfulError(ctx, "创建失败，原因："+err.Error())
			return
		}
	}
	helper.RestfulSuccess(ctx, "创建成功", nil)
}

// Show GET /dummy/{id}
func (r *AdminPermission) Show(ctx http.Context) {
	id := ctx.Request().Input("id")
	if id == "" {
		helper.RestfulError(ctx, "id:不能为空")
		return
	}
	var data models.AdminPermission
	if err := facades.Orm.Query().Where("id", id).First(&data); err != nil {
		helper.RestfulError(ctx, "查询错误1，原因: "+err.Error())
		return
	}
	if data.ID == 0 {
		helper.RestfulError(ctx, "查询失败：没有找到数据")
		return
	}
	var result models.PermissionResult
	result.ID = data.ID
	result.Name = data.Name
	result.Path = data.Path
	result.ParentId = data.ParentId
	result.Remark = data.Remark
	result.Redirect = data.Redirect
	result.Component = data.Component
	result.Active = data.Active
	result.Meta = helper.PermissionMetaToStruct(data.MetaJSON)
	result.Children, _ = (new(models.AdminPermission)).GetPermissionList(data.ID)
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"Item": result,
	})
}

// Edit GET /dummy/{id}/edit
func (r *AdminPermission) Edit(ctx http.Context) {
	var params requests.AdminAdminPermissionRequest
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "表单参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "表单参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	var permission models.AdminPermission
	id := ctx.Request().Input("id")
	if err := facades.Orm.Query().Where("id", id).FirstOrFail(&permission); err != nil {
		helper.RestfulError(ctx, "查询错误，原因: "+err.Error())
		return
	}

	if permission.ParentId != permission.ParentId {
		helper.RestfulError(ctx, "更新失败，原因：上级权限不能修改")
		return
	}

	// 更新数据
	update := &models.AdminPermission{
		Name:      params.Name,
		Path:      params.Path,
		ParentId:  params.ParentId,
		Remark:    params.Remark,
		Redirect:  params.Redirect,
		Component: params.Component,
		Active:    params.Active,
		MetaJSON:  helper.PermissionMetaToJSON(&params.Meta),
	}
	res, err := facades.Orm.Query().Where("id", id).Updates(&update)
	if err != nil {
		helper.RestfulError(ctx, "更新失败，原因："+err.Error())
		return
	}
	// 如果是按钮，添加到admin_permission_policy
	if params.Meta.Type == "button" {
		facades.Orm.Query().Where("path", params.Path).Updates(&models.AdminPermissionPolicy{
			Name: params.Name,
			Path: params.Path,
		})
	}
	if res.RowsAffected == 0 {
		helper.RestfulError(ctx, "更新失败，原因：未更新任何数据")
		return
	}
	helper.RestfulSuccess(ctx, "更新成功,影响行数："+strconv.FormatInt(res.RowsAffected, 10), nil)
}

// Delete DELETE /dummy
// @param id []int
func (r *AdminPermission) Delete(ctx http.Context) {
	ids := ctx.Request().QueryArray("id")
	if len(ids) == 0 {
		helper.RestfulError(ctx, "id:至少选择一条记录")
		return
	}
	var params []models.AdminPermission
	if err := facades.Orm.Query().Where("id", ids).Find(&params); err != nil {
		helper.RestfulError(ctx, "查询错误，原因: "+err.Error())
		return
	}
	// 判断是否有子级
	var errs []string
	var pn []string
	for _, param := range params {
		var count int64
		if err := facades.Orm.Query().Model(&r.Model).Where("parent_id", param.ID).Count(&count); err != nil {
			errs = append(errs, "查询错误，原因: "+err.Error())
		}
		if count > 0 {
			errs = append(errs, param.Name+"[ID: "+strconv.Itoa(int(param.ID))+"] 存在子级权限, 请先删除子级权限")
		}
		meta := helper.PermissionMetaToStruct(param.MetaJSON)
		if meta.Type == "button" {
			pn = append(pn, param.Path)
		}
	}
	if len(errs) > 0 {
		helper.RestfulError(ctx, errs[0])
		return
	}
	res, err := facades.Orm.Query().Delete(&r.Model, ids)
	if err != nil {
		helper.RestfulError(ctx, "删除失败："+err.Error())
		return
	}
	if len(pn) > 0 {
		facades.Orm.Query().Where("path in (?)", pn).Delete(&models.AdminPermissionPolicy{})
	}
	// 成功删除了 res.RowsAffected 条记录
	helper.RestfulSuccess(ctx, "成功删除了 "+strconv.Itoa(int(res.RowsAffected))+" 条记录", nil)
}
