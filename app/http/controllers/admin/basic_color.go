package admin

import (
	"Awesome/app/http/helper"
	"Awesome/app/http/requests"
	"Awesome/app/models"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"strconv"
)

type BasicColor struct {
	//Dependent services
}

func NewBasicColor() *BasicColor {
	return &BasicColor{
		//Inject services
	}
}

// Index is the index action of BasicColor controller.
func (r *BasicColor) Index(ctx http.Context) {
	name := ctx.Request().Query("name", "")
	code := ctx.Request().Query("code", "")
	page, _ := strconv.Atoi(ctx.Request().Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Request().Query("limit", "10"))
	// query
	query := facades.Orm.Query().Model(&models.BasicColor{})
	if name != "" { // 模糊
		query = query.Where("name like ?", "%"+name+"%")
	}
	if code != "" { // 模糊
		query = query.Where("code like ?", "%"+code+"%")
	}
	query.Order("id desc")
	// paginate
	var data []models.BasicColor
	var total int64
	if err := query.Paginate(page, limit, &data, &total); err != nil {
		helper.RestfulError(ctx, "查询失败，原因："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"Items": data,
		"Total": total,
	})
}

// Show is the show action of BasicColor controller.
func (r *BasicColor) Show(ctx http.Context) {
	key := ctx.Request().Input("key")
	if key == "" {
		helper.RestfulError(ctx, "表单参数错误：key不能为空")
		return
	}
	var data models.BasicColor
	if err := facades.Orm.Query().Model(&models.BasicColor{}).Where("id", key).OrWhere("name", key).First(&data); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"Item": data,
	})
}

// Create is the creation action of BasicColor controller.
func (r *BasicColor) Create(ctx http.Context) {
	var params requests.AdminBasicRequest
	if err := ctx.Request().Bind(&params); err != nil {
		helper.RestfulError(ctx, "表单参数错误："+err.Error())
		return
	}
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "表单参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "表单参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	// unique
	var count int64
	facades.Orm.Query().Model(&models.BasicColor{}).Where("name", params.Name).Count(&count)
	if count > 0 {
		helper.RestfulError(ctx, "创建失败：名称已存在")
		return
	}
	// Create your model here
	data := &models.BasicColor{
		Name:   params.Name,
		Code:   params.Code,
		Remark: params.Remark,
	}
	if err := facades.Orm.Query().Create(data); err != nil {
		helper.RestfulError(ctx, "创建失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "创建成功", nil)
}

// Store is the store action of BasicColor controller.
func (r *BasicColor) Store(ctx http.Context) {
	key := ctx.Request().Input("key")
	if key == "" {
		helper.RestfulError(ctx, "表单参数错误：key不能为空")
		return
	}
	var params requests.AdminBasicRequest
	if err := ctx.Request().Bind(&params); err != nil {
		helper.RestfulError(ctx, "表单参数错误："+err.Error())
		return
	}
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "表单参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "表单参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	var data models.BasicColor
	if err := facades.Orm.Query().Model(&models.BasicColor{}).Where("id", key).OrWhere("name", key).First(&data); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	update := &models.BasicColor{
		Name:   data.Name,
		Code:   data.Code,
		Remark: data.Remark,
	}
	if err := facades.Orm.Query().Model(&models.BasicColor{}).Where("id", data.ID).Updates(update); err != nil {
		helper.RestfulError(ctx, "更新失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "更新成功", nil)
}

// Delete is the delete action of BasicColor controller.
func (r *BasicColor) Delete(ctx http.Context) {
	ids := ctx.Request().QueryArray("ids")
	if len(ids) == 0 {
		helper.RestfulError(ctx, "表单参数错误：ids不能为空")
		return
	}
	if err := facades.Orm.Query().Delete(&models.BasicColor{}, ids); err != nil {
		helper.RestfulError(ctx, "删除失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "删除成功", nil)
}
