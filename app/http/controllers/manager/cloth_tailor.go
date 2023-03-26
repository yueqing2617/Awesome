package manager

import (
	"Awesome/app/http/helper"
	"Awesome/app/models"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"strconv"
)

type ClothTailor struct {
	//Dependent services
	Model models.ClothTailor
}

func NewClothTailor() *ClothTailor {
	return &ClothTailor{
		//Inject services
	}
}

// Index GET /dummy
func (r *ClothTailor) Index(ctx http.Context) {
	page, _ := strconv.Atoi(ctx.Request().Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Request().Query("limit", "10"))
	orderCode := ctx.Request().Query("cloth_order_code", "")
	clothStyleCode := ctx.Request().Query("cloth_style_code", "")
	order := ctx.Request().Query("order", "id desc")
	// query
	query := facades.Orm.Query().Model(&r.Model)
	if orderCode != "" {
		query = query.Where("cloth_order_code", "like", "%"+orderCode+"%")
	}
	if clothStyleCode != "" {
		query = query.Where("cloth_style_code", "like", "%"+clothStyleCode+"%")
	}
	// 排序
	query = query.Order(order)
	// 分页
	var total int64
	var list []models.ClothTailor
	if err := query.With("ClothOrder").Paginate(page, limit, &list, &total); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"total": total,
		"items": list,
	})
}

// Create POST /dummy
func (r *ClothTailor) Create(ctx http.Context) {

}

// Show GET /dummy/{id}
func (r *ClothTailor) Show(ctx http.Context) {
}

// Edit GET /dummy/{id}/edit
func (r *ClothTailor) Edit(ctx http.Context) {
}

// Delete DELETE /dummy
// @param id []int
func (r *ClothTailor) Delete(ctx http.Context) {
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
