package manager

import (
	"Awesome/app/http/helper"
	"Awesome/app/http/requests"
	"Awesome/app/models"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"strconv"
)

type Admin struct {
	//Dependent services
	Model models.Admin
}

func NewAdmin() *Admin {
	return &Admin{
		//Inject services
	}
}

// Index GET /dummy
func (r *Admin) Index(ctx http.Context) {
	page, _ := strconv.Atoi(ctx.Request().Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Request().Query("limit", "10"))
	phone := ctx.Request().Query("phone", "")
	roleID, _ := strconv.Atoi(ctx.Request().Query("role_id", "0"))
	nickName := ctx.Request().Query("nickname", "")
	order := ctx.Request().Query("order", "id desc")
	// 查询条件
	query := facades.Orm.Query().Model(&r.Model)
	if phone != "" {
		query = query.Where("phone like ?", "%"+phone+"%")
	}
	if roleID != 0 {
		query = query.Where("role_id = ?", roleID)
	}
	if nickName != "" {
		query = query.Where("nickname like ?", "%"+nickName+"%")
	}
	query.Order(order)
	// 查询数据
	type AdminResult struct {
		ID        uint              `json:"id"`
		Phone     string            `json:"phone"`
		Nickname  string            `json:"nickname"`
		RoleID    uint              `json:"role_id"`
		Role      *models.AdminRole `json:"role"`
		Avatar    string            `json:"avatar"`
		Email     string            `json:"email"`
		Gender    string            `json:"gender"`
		CreatedAt int64             `json:"created_at"`
		UpdatedAt int64             `json:"updated_at"`
	}
	var total int64
	var data []AdminResult
	if err := query.With("Role").Paginate(page, limit, &data, &total); err != nil {
		helper.RestfulError(ctx, "查询错误，原因: "+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"items": data,
		"total": total,
	})
}

// Create POST /dummy
func (r *Admin) Create(ctx http.Context) {
	var params requests.AdminAdminRequest
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "表单参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "表单参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	pwd, _ := facades.Hash.Make(params.Password)
	data := &models.Admin{
		Nickname: params.Nickname,
		Phone:    params.Phone,
		Password: pwd,
		Avatar:   params.Avatar,
		Email:    params.Email,
		Remark:   params.Remark,
		RoleID:   params.RoleID,
		Role:     nil,
		Gender:   params.Gender,
	}
	if err := facades.Orm.Query().Create(data); err != nil {
		helper.RestfulError(ctx, "创建失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "创建成功", nil)
}

// Show GET /dummy/{id}
func (r *Admin) Show(ctx http.Context) {
	id := ctx.Request().Input("id")
	if id == "" {
		helper.RestfulError(ctx, "id:不能为空")
		return
	}
	var data struct {
		ID        uint              `json:"id"`
		Nickname  string            `json:"nickname"`
		Phone     string            `json:"phone"`
		Avatar    string            `json:"avatar"`
		Email     string            `json:"email"`
		Remark    string            `json:"remark"`
		RoleID    uint              `json:"role_id"`
		Role      *models.AdminRole `json:"role"`
		CreatedAt int64             `json:"created_at"`
		UpdatedAt int64             `json:"updated_at"`
	}
	if err := facades.Orm.Query().Model(&r.Model).Where("id = ?", id).With("Role").FirstOrFail(&data); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	if data.ID == 0 {
		helper.RestfulError(ctx, "查询失败：记录不存在")
		return
	}
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"item": data,
	})
}

// Edit GET /dummy/{id}/edit
func (r *Admin) Edit(ctx http.Context) {
	var params requests.AdminAdminRequest
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
		helper.RestfulError(ctx, "id:不能为空")
		return
	}
	var data models.Admin
	if err := facades.Orm.Query().Model(&r.Model).Where("id = ?", id).FirstOrFail(&data); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	if data.ID == 0 {
		helper.RestfulError(ctx, "查询失败：记录不存在")
		return
	}
	update := map[string]interface{}{
		"nickname": params.Nickname,
		"avatar":   params.Avatar,
		"email":    params.Email,
		"remark":   params.Remark,
		"role_id":  params.RoleID,
		"gender":   params.Gender,
	}
	if err := facades.Orm.Query().Model(&r.Model).Where("id = ?", data.ID).Save(&update); err != nil {
		helper.RestfulError(ctx, "更新失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "更新成功", nil)
}

// Delete DELETE /dummy
// @param id []int
func (r *Admin) Delete(ctx http.Context) {
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
