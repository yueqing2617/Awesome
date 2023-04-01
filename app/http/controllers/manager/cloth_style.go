package manager

import (
	"Awesome/app/http/helper"
	"Awesome/app/http/requests"
	"Awesome/app/models"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"strconv"
)

type ClothStyle struct {
	//Dependent services
	Model models.ClothStyle
}

func NewClothStyle() *ClothStyle {
	return &ClothStyle{
		//Inject services
	}
}

// Index GET /dummy
func (r *ClothStyle) Index(ctx http.Context) {
	page, _ := strconv.Atoi(ctx.Request().Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Request().Query("limit", "10"))
	name := ctx.Request().Query("name", "")
	code := ctx.Request().Query("code", "")
	year := ctx.Request().Query("year", "")
	season := ctx.Request().Query("season", "")
	order := ctx.Request().Query("order", "id desc")
	// 查询
	query := facades.Orm.Query().Model(&r.Model)
	if name != "" {
		query = query.Where("name", "like", "%"+name+"%")
	}
	if code != "" {
		query = query.Where("code", "like", "%"+code+"%")
	}
	if year != "" {
		query = query.Where("year", "like", "%"+year+"%")
	}
	if season != "" {
		query = query.Where("season", "like", "%"+season+"%")
	}
	// 排序
	query = query.Order(order)
	// 分页
	var total int64
	var data []models.ClothStyle
	if err := query.Paginate(page, limit, &data, &total); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	result := (new(models.ClothStyle)).Value(&data)
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"items": result,
		"total": total,
	})
}

// Create POST /dummy
func (r *ClothStyle) Create(ctx http.Context) {
	var params requests.AdminClothStyleRequest
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "请求参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "请求参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	data := &models.ClothStyle{
		Name:         params.Name,
		Code:         params.Code,
		Picture:      params.Picture,
		Colors:       helper.ArrayToString(params.Colors),
		Sizes:        helper.ArrayToString(params.Sizes),
		Year:         params.Year,
		Season:       params.Season,
		UnitPrice:    params.UnitPrice,
		ProcedureStr: helper.ProceduresToJson(&params.Procedures),
		Remark:       params.Remark,
	}
	if err := facades.Orm.Query().Model(&r.Model).Create(data); err != nil {
		helper.RestfulError(ctx, "创建失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "创建成功", nil)
}

// Show GET /dummy/{id}
func (r *ClothStyle) Show(ctx http.Context) {
	id := ctx.Request().Input("id")
	if id == "" {
		helper.RestfulError(ctx, "id:不能为空")
		return
	}
	var data models.ClothStyle
	if err := facades.Orm.Query().Model(&r.Model).Where("id", id).First(&data); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	result := models.ClothStyleResult{
		ID:         data.ID,
		Name:       data.Name,
		Code:       data.Code,
		Picture:    data.Picture,
		Colors:     helper.StringToArray(data.Colors),
		Sizes:      helper.StringToArray(data.Sizes),
		Year:       data.Year,
		Season:     data.Season,
		UnitPrice:  data.UnitPrice,
		Procedures: helper.JsonToProcedures(data.ProcedureStr),
		Remark:     data.Remark,
	}
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"item": result,
	})
}

// Edit GET /dummy/{id}/edit
func (r *ClothStyle) Edit(ctx http.Context) {
	var params requests.AdminClothStyleRequest
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
	if id == "" {
		helper.RestfulError(ctx, "id:不能为空")
		return
	}
	data := &models.ClothStyle{
		Name:         params.Name,
		Picture:      params.Picture,
		Colors:       helper.ArrayToString(params.Colors),
		Sizes:        helper.ArrayToString(params.Sizes),
		Year:         params.Year,
		Season:       params.Season,
		UnitPrice:    params.UnitPrice,
		ProcedureStr: helper.ProceduresToJson(&params.Procedures),
		Remark:       params.Remark,
	}
	res, err := facades.Orm.Query().Model(&r.Model).Where("id", id).Updates(data)
	if err != nil {
		helper.RestfulError(ctx, "更新失败："+err.Error())
		return
	}
	if res.RowsAffected == 0 {
		helper.RestfulError(ctx, "更新失败：未更新任何数据")
		return
	}
	helper.RestfulSuccess(ctx, "更新成功,更新了 "+strconv.Itoa(int(res.RowsAffected))+" 条记录", nil)
}

// Delete DELETE /dummy
// @param id []int
func (r *ClothStyle) Delete(ctx http.Context) {
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
