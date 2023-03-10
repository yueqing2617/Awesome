package admin

import (
	"Awesome/app/http/helper"
	"Awesome/app/http/requests"
	"Awesome/app/models"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"strconv"
)

type BasicProcedure struct {
	//Dependent services
}

func NewBasicProcedure() *BasicProcedure {
	return &BasicProcedure{
		//Inject services
	}
}

// Index is the index action of BasicProcedure controller.
func (r *BasicProcedure) Index(ctx http.Context) {
	name := ctx.Request().Query("name", "")
	code := ctx.Request().Query("code", "")
	page, _ := strconv.Atoi(ctx.Request().Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Request().Query("limit", "10"))
	// query
	query := facades.Orm.Query().Model(&models.BasicProcedure{})
	if name != "" { // 模糊
		query = query.Where("name like ?", "%"+name+"%")
	}
	if code != "" { // 模糊
		query = query.Where("code like ?", "%"+code+"%")
	}
	query.Order("id desc")
	// paginate
	var data []models.BasicProcedure
	var total int64
	if err := query.Paginate(page, limit, &data, &total); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"Items": data,
		"Total": total,
	})
}

// Show is the show action of BasicProcedure controller.
func (r *BasicProcedure) Show(ctx http.Context) {
	key := ctx.Request().Input("key")
	if key == "" {
		helper.RestfulError(ctx, "表单参数错误：key不能为空")
		return
	}
	var data models.BasicSize
	if err := facades.Orm.Query().Model(&models.BasicProcedure{}).Where("id", key).OrWhere("name", key).First(&data); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"Item": data,
	})
}

// Create is the creation action of BasicProcedure controller.
func (r *BasicProcedure) Create(ctx http.Context) {
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
	facades.Orm.Query().Model(&models.BasicProcedure{}).Where("name", params.Name).Count(&count)
	if count > 0 {
		helper.RestfulError(ctx, "创建失败：名称已存在")
		return
	}
	// Create your model here
	data := models.BasicProcedure{
		Name:   params.Name,
		Code:   params.Code,
		Remark: params.Remark,
	}
	if err := facades.Orm.Query().Create(&data); err != nil {
		helper.RestfulError(ctx, "创建失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "创建成功", nil)
}

// Store is the store action of BasicProcedure controller.
func (r *BasicProcedure) Store(ctx http.Context) {
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
	// query
	var data models.BasicProcedure
	if err := facades.Orm.Query().Model(&models.BasicProcedure{}).Where("id", key).OrWhere("name", key).First(&data); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	// update
	update := map[string]interface{}{
		"name":   params.Name,
		"code":   params.Code,
		"remark": params.Remark,
	}
	if err := facades.Orm.Query().Model(&models.BasicProcedure{}).Where("id", data.ID).Updates(&update); err != nil {
		helper.RestfulError(ctx, "更新失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "更新成功", nil)
}

// Delete is the deletion action of BasicProcedure controller.
func (r *BasicProcedure) Delete(ctx http.Context) {
	ids := ctx.Request().QueryArray("ids")
	if len(ids) == 0 {
		helper.RestfulError(ctx, "表单参数错误：ids不能为空")
		return
	}
	if err := facades.Orm.Query().Delete(&models.BasicProcedure{}, ids); err != nil {
		helper.RestfulError(ctx, "删除失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "删除成功", nil)
}
