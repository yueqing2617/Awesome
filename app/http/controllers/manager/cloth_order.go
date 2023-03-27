package manager

import (
	"Awesome/app/events"
	"Awesome/app/http/helper"
	"Awesome/app/http/requests"
	"Awesome/app/models"
	"github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"strconv"
)

type ClothOrder struct {
	//Dependent services
	Model models.ClothOrder
}

func NewClothOrder() *ClothOrder {
	return &ClothOrder{
		//Inject services
	}
}

// Index GET /dummy
func (r *ClothOrder) Index(ctx http.Context) {
	page, _ := strconv.Atoi(ctx.Request().Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Request().Query("limit", "10"))
	name := ctx.Request().Query("name", "")
	code := ctx.Request().Query("code", "")
	customerId := ctx.Request().Query("customer_id", "")
	salesman := ctx.Request().Query("salesman_id", "")
	order := ctx.Request().Query("order", "id desc")
	query := facades.Orm.Query().Model(&r.Model)
	if name != "" {
		query = query.Where("name", "like", "%"+name+"%").OrWhere("cloth_style_name like ?", "%"+name+"%")
	}
	if code != "" {
		query = query.Where("code", "like", "%"+code+"%").OrWhere("cloth_style_code like ?", "%"+code+"%")
	}
	if customerId != "" {
		query = query.Where("customer_id", customerId)
	}
	if salesman != "" {
		query = query.Where("salesman_id", salesman)
	}
	// 排序
	query = query.Order(order)
	// 分页
	var total int64
	var list []models.ClothOrder
	if err := query.With("Customer").With("Salesman").Paginate(page, limit, &list, &total); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	var result []models.ClothOrderResult
	for _, item := range list {
		result = append(result, models.ClothOrderResult{
			ID:                  item.ID,
			Name:                item.Name,
			Code:                item.Code,
			CustomerID:          item.CustomerID,
			Customer:            item.Customer,
			DeliveryDate:        item.DeliveryDate,
			OrderType:           item.OrderType,
			ClothStyleCode:      item.ClothStyleCode,
			ClothStyleName:      item.ClothStyleName,
			ClothStyleColors:    helper.StringToArray(item.ClothStyleColors),
			ClothStyleSizes:     helper.StringToArray(item.ClothStyleSizes),
			ClothStyleYear:      item.ClothStyleYear,
			ClothStyleSeason:    item.ClothStyleSeason,
			ClothStyleUnitPrice: item.ClothStyleUnitPrice,
			Total:               item.Total,
			TotalPrice:          item.TotalPrice,
			Contains:            helper.JSONToContains(item.ContainsStr),
			Procedures:          helper.JsonToProcedures(item.ProceduresStr),
			Remark:              item.Remark,
		})
	}
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"items": result,
		"total": total,
	})
}

// Create POST /dummy
func (r *ClothOrder) Create(ctx http.Context) {
	var params requests.AdminClothOrderRequest
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "请求参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "请求参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	var data = &models.ClothOrder{
		Name:                params.Name,
		Code:                params.Code,
		CustomerID:          params.CustomerID,
		DeliveryDate:        params.DeliveryDate,
		OrderType:           params.OrderType,
		SalesmanID:          params.SalesmanID,
		ClothStyleCode:      params.ClothStyleCode,
		ClothStyleName:      params.ClothStyleName,
		ClothStylePicture:   params.ClothStylePicture,
		ClothStyleColors:    helper.ArrayToString(params.ClothStyleColors),
		ClothStyleSizes:     helper.ArrayToString(params.ClothStyleSizes),
		ClothStyleYear:      params.ClothStyleYear,
		ClothStyleSeason:    params.ClothStyleSeason,
		ClothStyleUnitPrice: params.ClothStyleUnitPrice,
		Total:               params.Total,
		TotalPrice:          params.TotalPrice,
		ContainsStr:         helper.ContainsToJSON(&params.Contains),
		ProceduresStr:       helper.ProceduresToJson(&params.Procedures),
		Status:              1,
		Remark:              params.Remark,
	}
	if err := facades.Orm.Query().Model(&r.Model).Create(data); err != nil {
		helper.RestfulError(ctx, "创建失败："+err.Error())
		return
	}
	if err := facades.Orm.Query().UpdateOrCreate(&models.ClothStyle{}, &models.ClothStyle{Code: data.ClothStyleCode}, &models.ClothStyle{
		Name:         data.ClothStyleName,
		Code:         data.ClothStyleCode,
		Picture:      data.ClothStylePicture,
		Colors:       data.ClothStyleColors,
		Sizes:        data.ClothStyleSizes,
		Year:         data.ClothStyleYear,
		Season:       data.ClothStyleSeason,
		UnitPrice:    data.ClothStyleUnitPrice,
		ProcedureStr: helper.ProceduresToJson(&params.Procedures),
	}); err != nil {
		helper.RestfulError(ctx, "更新失败："+err.Error())
		return
	}
	customer := models.BasicCustomer{}
	if err := facades.Orm.Query().Model(&customer).Where("id", data.CustomerID).First(&customer); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	_ = facades.Event.Job(&events.OrderCreated{}, []event.Arg{
		{Type: "string", Value: data.Code},
		{Type: "string", Value: data.ClothStyleCode},
		{Type: "uint", Value: data.Total},
		{Type: "string", Value: "create"},
		{Type: "string", Value: customer.Name},
	}).Dispatch()
	helper.RestfulSuccess(ctx, "创建成功", nil)
}
func (r *ClothOrder) Show(ctx http.Context) {
	id := ctx.Request().Input("id")
	if id == "" {
		helper.RestfulError(ctx, "id:不能为空")
		return
	}
	var data models.ClothOrder
	if err := facades.Orm.Query().Model(&r.Model).With("Customer").Where("id", id).First(&data); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}

	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"item": http.Json{
			"id":                     data.ID,
			"name":                   data.Name,
			"code":                   data.Code,
			"customer":               data.Customer,
			"delivery_date":          data.DeliveryDate,
			"order_type":             data.OrderType,
			"salesman_id":            data.SalesmanID,
			"salesman":               data.Salesman,
			"cloth_style_code":       data.ClothStyleCode,
			"cloth_style_name":       data.ClothStyleName,
			"cloth_style_picture":    data.ClothStylePicture,
			"cloth_style_colors":     helper.StringToArray(data.ClothStyleColors),
			"cloth_style_sizes":      helper.StringToArray(data.ClothStyleSizes),
			"cloth_style_year":       data.ClothStyleYear,
			"cloth_style_season":     data.ClothStyleSeason,
			"cloth_style_unit_price": data.ClothStyleUnitPrice,
			"total":                  data.Total,
			"total_price":            data.TotalPrice,
			"contains":               helper.JSONToContains(data.ContainsStr),
			"procedures":             helper.JsonToProcedures(data.ProceduresStr),
			"remark":                 data.Remark,
		},
	})
}

// Edit GET /dummy/{id}/edit
func (r *ClothOrder) Edit(ctx http.Context) {
	var params requests.AdminClothOrderRequest
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "请求参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "请求参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	id := ctx.Request().Input("id")
	// 更新数据
	var data = &models.ClothOrder{
		Name:                params.Name,
		CustomerID:          params.CustomerID,
		DeliveryDate:        params.DeliveryDate,
		OrderType:           params.OrderType,
		SalesmanID:          params.SalesmanID,
		ClothStyleCode:      params.ClothStyleCode,
		ClothStyleName:      params.ClothStyleName,
		ClothStylePicture:   params.ClothStylePicture,
		ClothStyleColors:    helper.ArrayToString(params.ClothStyleColors),
		ClothStyleSizes:     helper.ArrayToString(params.ClothStyleSizes),
		ClothStyleYear:      params.ClothStyleYear,
		ClothStyleSeason:    params.ClothStyleSeason,
		ClothStyleUnitPrice: params.ClothStyleUnitPrice,
		Total:               params.Total,
		TotalPrice:          params.TotalPrice,
		ContainsStr:         helper.ContainsToJSON(&params.Contains),
		ProceduresStr:       helper.ProceduresToJson(&params.Procedures),
		Remark:              params.Remark,
	}
	res, err := facades.Orm.Query().Model(&r.Model).Where("id", id).Updates(&data)
	if err != nil {
		helper.RestfulError(ctx, "更新失败："+err.Error())
		return
	}
	err = facades.Orm.Query().UpdateOrCreate(&models.ClothStyle{}, &models.ClothStyle{Code: data.ClothStyleCode}, &models.ClothStyle{
		Name:         data.ClothStyleName,
		Code:         data.ClothStyleCode,
		Picture:      data.ClothStylePicture,
		Colors:       data.ClothStyleColors,
		Sizes:        data.ClothStyleSizes,
		Year:         data.ClothStyleYear,
		Season:       data.ClothStyleSeason,
		UnitPrice:    data.ClothStyleUnitPrice,
		ProcedureStr: helper.ProceduresToJson(&params.Procedures),
	})
	// 成功更新了 res.RowsAffected 条记录
	if res.RowsAffected == 0 {
		helper.RestfulError(ctx, "更新失败：没有更新任何数据")
		return
	}
	customer := models.BasicCustomer{}
	facades.Orm.Query().Model(&customer).Where("id", data.CustomerID).First(&customer)
	_ = facades.Event.Job(&events.OrderCreated{}, []event.Arg{
		{Type: "string", Value: data.Code},
		{Type: "string", Value: data.ClothStyleColors},
		{Type: "uint", Value: data.Total},
		{Type: "string", Value: "update"},
		{Type: "string", Value: data.ClothStyleName},
	}).Dispatch()

	helper.RestfulSuccess(ctx, "更新成功", nil)
}

// Delete DELETE /dummy
// @param id []int
func (r *ClothOrder) Delete(ctx http.Context) {
	ids := ctx.Request().QueryArray("id")
	if len(ids) == 0 {
		helper.RestfulError(ctx, "ids:至少选择一条记录")
	}
	res, err := facades.Orm.Query().Delete(&r.Model, ids)
	if err != nil {
		helper.RestfulError(ctx, "删除失败："+err.Error())
		return
	}
	// 成功删除了 res.RowsAffected 条记录
	helper.RestfulSuccess(ctx, "成功删除了 "+strconv.Itoa(int(res.RowsAffected))+" 条记录", nil)
}
