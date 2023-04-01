package manager

import (
	"Awesome/app/http/helper"
	"Awesome/app/http/requests"
	"Awesome/app/models"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"strconv"
)

type BasicCustomer struct {
	//Dependent services
	Model models.BasicCustomer
}

func NewBasicCustomer() *BasicCustomer {
	return &BasicCustomer{
		//Inject services
	}
}

// Index GET /dummy
func (r *BasicCustomer) Index(ctx http.Context) {
	page, _ := strconv.Atoi(ctx.Request().Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Request().Query("limit", "10"))
	name := ctx.Request().Query("name", "")
	code := ctx.Request().Query("code", "")
	phone := ctx.Request().Query("phone", "")
	order := ctx.Request().Query("order", "id desc")
	// 查询
	query := facades.Orm.Query().Model(r.Model)
	if name != "" {
		query = query.Where("name like ?", "%"+name+"%")
	}
	if code != "" {
		query = query.Where("code like ?", "%"+code+"%")
	}
	if phone != "" {
		query = query.Where("phone like ?", "%"+phone+"%")
	}
	query.Order(order)
	// 分页
	var data []models.BasicCustomer
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
func (r *BasicCustomer) Create(ctx http.Context) {
	var params requests.AdminBasicCustomerRequest
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "表单参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "表单参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	var data models.BasicCustomer
	err = facades.Orm.Query().Model(r.Model).Where("name = ?", params.Name).FirstOrCreate(&data, models.BasicCustomer{
		Name:    params.Name,
		Code:    params.Code,
		Phone:   params.Phone,
		Address: params.Address,
		Gender:  params.Gender,
		Company: params.Company,
		Remark:  params.Remark,
	})
	if err != nil {
		helper.RestfulError(ctx, "创建失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "创建成功", nil)
}

// Show GET /dummy/{id}
func (r *BasicCustomer) Show(ctx http.Context) {
	key := ctx.Request().Input("key")
	if key == "" {
		helper.RestfulError(ctx, "表单参数错误：key不能为空")
		return
	}
	var data models.BasicCustomer
	err := facades.Orm.Query().Model(r.Model).Where("id = ?", key).First(&data)
	if err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	if data.ID == 0 {
		helper.RestfulError(ctx, "查询失败：没有找到数据")
		return
	}
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"Item": data,
	})
}

// Edit GET /dummy/{id}/edit
func (r *BasicCustomer) Edit(ctx http.Context) {
	var params requests.AdminBasicCustomerRequest
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "表单参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "表单参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	key := ctx.Request().Input("id")
	if key == "" {
		helper.RestfulError(ctx, "表单参数错误：id不能为空")
		return
	}
	var data models.BasicCustomer
	if err := facades.Orm.Query().Model(r.Model).Where("id = ?", key).FirstOrFail(&data); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	data.Name = params.Name
	data.Code = params.Code
	data.Phone = params.Phone
	data.Address = params.Address
	data.Company = params.Company
	data.Gender = params.Gender
	data.Remark = params.Remark
	if err := facades.Orm.Query().Model(&r.Model).Where("id = ?", key).Save(&data); err != nil {
		helper.RestfulError(ctx, "更新失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "更新成功", nil)
}

// Delete DELETE /dummy
// @param id []int
func (r *BasicCustomer) Delete(ctx http.Context) {
	ids := ctx.Request().QueryArray("id")
	if len(ids) == 0 {
		helper.RestfulError(ctx, "id:至少选择一条记录")
		return
	}
	res, err := facades.Orm.Query().Delete(&r.Model, ids)
	if err != nil {
		helper.RestfulError(ctx, "删除失败："+err.Error())
		return
	}
	// 成功删除了 res.RowsAffected 条记录
	helper.RestfulSuccess(ctx, "成功删除了 "+strconv.Itoa(int(res.RowsAffected))+" 条记录", nil)
}
