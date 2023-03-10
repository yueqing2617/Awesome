package admin

import (
	"Awesome/app/http/helper"
	"Awesome/app/http/requests"
	"Awesome/app/models"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type BasicCustomer struct {
	//Dependent services
}

func NewBasicCustomer() *BasicCustomer {
	return &BasicCustomer{
		//Inject services
	}
}

// Index is the index action of BasicCustomer controller.
func (r *BasicCustomer) Index(ctx http.Context) {
	name := ctx.Request().Query("name", "")
	code := ctx.Request().Query("code", "")
	phone := ctx.Request().Query("phone", "")
	gender := ctx.Request().Query("gender", "")
	// query
	query := facades.Orm.Query().Model(&models.BasicCustomer{})
	if name != "" { // 模糊
		query = query.Where("name like ?", "%"+name+"%")
	}
	if code != "" { // 模糊
		query = query.Where("code like ?", "%"+code+"%")
	}
	if phone != "" { // 模糊
		query = query.Where("phone like ?", "%"+phone+"%")
	}
	if gender != "" {
		query = query.Where("gender = ?", gender)
	}
	query.Order("id desc")
	// paginate
	var data []models.BasicCustomer
	var total int64
	if err := query.Paginate(1, 10, &data, &total); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"Items": data,
		"Total": total,
	})
}

// Show is the show action of BasicCustomer controller.
func (r *BasicCustomer) Show(ctx http.Context) {
	key := ctx.Request().Input("key")
	if key == "" {
		helper.RestfulError(ctx, "表单参数错误：key不能为空")
		return
	}
	var data models.BasicCustomer
	if err := facades.Orm.Query().Model(&models.BasicCustomer{}).Where("id", key).OrWhere("name", key).First(&data); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "查询成功", data)
}

// Create is the creation action of BasicCustomer controller.
func (r *BasicCustomer) Create(ctx http.Context) {
	var params requests.AdminBasicCustomerRequest
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
	facades.Orm.Query().Model(&models.BasicSize{}).Where("phone", params.Phone).Count(&count)
	if count > 0 {
		helper.RestfulError(ctx, "创建失败：手机号已存在")
		return
	}
	// Create your model here
	data := models.BasicCustomer{
		Name:    params.Name,
		Code:    params.Code,
		Phone:   params.Phone,
		Remark:  params.Remark,
		Gender:  params.Gender,
		Company: params.Company,
		Address: params.Address,
	}
	if err := facades.Orm.Query().Create(&data); err != nil {
		helper.RestfulError(ctx, "创建失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "创建成功", data)
}

// Store is the store action of BasicCustomer controller.
func (r *BasicCustomer) Store(ctx http.Context) {
	key := ctx.Request().Input("key")
	if key == "" {
		helper.RestfulError(ctx, "表单参数错误：key不能为空")
		return
	}
	var params requests.AdminBasicCustomerRequest
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
	var data models.BasicCustomer
	if err := facades.Orm.Query().Model(&data).Where("id", key).OrWhere("name", key).First(&data); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	// update
	update := map[string]interface{}{
		"name":    params.Name,
		"code":    params.Code,
		"remark":  params.Remark,
		"gender":  params.Gender,
		"company": params.Company,
		"address": params.Address,
		"phone":   params.Phone,
	}
	if err := facades.Orm.Query().Model(&data).Updates(update); err != nil {
		helper.RestfulError(ctx, "更新失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "更新成功", nil)
}

// Delete is the deletion action of BasicCustomer controller.
func (r *BasicCustomer) Delete(ctx http.Context) {
	ids := ctx.Request().QueryArray("ids")
	if len(ids) == 0 {
		helper.RestfulError(ctx, "表单参数错误：ids不能为空")
		return
	}
	if err := facades.Orm.Query().Delete(&models.BasicCustomer{}, ids); err != nil {
		helper.RestfulError(ctx, "删除失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "删除成功", nil)
}
