package admin

import (
	"Awesome/app/http/helper"
	"Awesome/app/http/requests"
	"Awesome/app/models"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"strconv"
)

type Admins struct {
	//Dependent services
}

func NewAdmins() *Admins {
	return &Admins{
		//Inject services
	}
}

// Index is the index action of Admins controller.
func (_t *Admins) Index(ctx http.Context) {
	phone := ctx.Request().Query("phone", "")
	nickname := ctx.Request().Query("nickname", "")
	email := ctx.Request().Query("email", "")
	roleName := ctx.Request().Query("role_name", "")
	page, _ := strconv.Atoi(ctx.Request().Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Request().Query("limit", "10"))
	createdStart := ctx.Request().Query("created_start", "")
	createdEnd := ctx.Request().Query("created_end", "")
	// query
	query := facades.Orm.Query().Model(&models.Admin{})
	if phone != "" {
		query = query.Where("phone", phone)
	}
	if nickname != "" {
		query = query.Where("nickname like ?", "%"+nickname+"%")
	}
	if email != "" { // 模糊
		query = query.Where("email like ?", "%"+email+"%")
	}
	if roleName != "" {
		query = query.Where("role_name", roleName)
	}
	if createdStart != "" {
		query = query.Where("created_at", ">=", createdStart)
	}
	if createdEnd != "" {
		query = query.Where("created_at", "<=", createdEnd)
	}
	type result struct {
		ID        uint
		Nickname  string
		Phone     string
		Email     string
		RoleName  string
		Gender    string
		Avatar    string
		Role      interface{}
		Remark    string
		CreatedAt string
		UpdatedAt string
	}
	// count
	var count int64
	var admins []result
	if err := query.Paginate(page, limit, &admins, &count); err != nil {
		helper.RestfulError(ctx, "服务器错误："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "请求成功", http.Json{"Items": admins, "Total": count})
}

// Create is the creation action of Admins controller.
func (_t *Admins) Create(ctx http.Context) {
	var params = requests.AdminAdminCreateRequest{}
	if err := ctx.Request().Bind(&params); err != nil {
		helper.RestfulError(ctx, "请求参数错误："+err.Error())
		return
	}
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "请求参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "请求参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	// Create your model here
	data := &models.Admin{
		Phone:    params.Phone,
		Password: helper.PasswordEncrypt(params.Password),
		Nickname: params.Nickname,
		Gender:   params.Gender,
		Email:    params.Email,
		Avatar:   params.Avatar,
		RoleName: params.RoleName,
		Remark:   params.Remark,
	}
	if err = facades.Orm.Query().Create(&data); err != nil {
		helper.RestfulError(ctx, "服务器错误："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "新增成功", nil)
}

// Store is the storage action of Admins controller.
func (_t *Admins) Store(ctx http.Context) {
	id := ctx.Request().Input("id")
	if id == "" {
		helper.RestfulError(ctx, "请求参数错误: 缺少id")
		return
	}
	data := &models.Admin{}
	if err := facades.Orm.Query().Where("id", id).First(&data); err != nil {
		helper.RestfulError(ctx, "服务器错误："+err.Error())
		return
	}
	if data.ID == 0 {
		helper.RestfulError(ctx, "请求参数错误: 该数据不存在")
		return
	}
	var params = requests.AdminAdminEditRequest{}
	if err := ctx.Request().Bind(&params); err != nil {
		helper.RestfulError(ctx, "请求参数错误："+err.Error())
		return
	}
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "请求参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "请求参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	// Update your model here
	data.Nickname = params.Nickname
	data.Email = params.Email
	data.Avatar = params.Avatar
	data.RoleName = params.RoleName
	data.Remark = params.Remark
	data.Gender = params.Gender
	if err = facades.Orm.Query().Save(&data); err != nil {
		helper.RestfulError(ctx, "服务器错误："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "更新成功", nil)
}

// Show is the show action of Admins controller.
func (_t *Admins) Show(ctx http.Context) {
	id := ctx.Request().Input("id")
	if id == "" {
		helper.RestfulError(ctx, "请求参数错误: 缺少id")
		return
	}
	data := &models.Admin{}
	if err := facades.Orm.Query().Where("id", id).First(&data); err != nil {
		helper.RestfulError(ctx, "服务器错误："+err.Error())
		return
	}
	if data.ID == 0 {
		helper.RestfulError(ctx, "请求参数错误: 该数据不存在")
		return
	}
	helper.RestfulSuccess(ctx, "请求成功", http.Json{
		"Item": http.Json{
			"id":         data.ID,
			"phone":      data.Phone,
			"nickname":   data.Nickname,
			"gender":     data.Gender,
			"email":      data.Email,
			"avatar":     data.Avatar,
			"role_name":  data.RoleName,
			"role":       data.Role,
			"remark":     data.Remark,
			"created_at": data.CreatedAt,
			"updated_at": data.UpdatedAt,
		},
	})
}

// Delete is the deletion action of Admins controller.
func (_t *Admins) Delete(ctx http.Context) {
	ids := ctx.Request().QueryArray("ids")
	if len(ids) == 0 {
		helper.RestfulError(ctx, "请求参数错误: 缺少ids")
		return
	}
	if err := facades.Orm.Query().Delete(&models.Admin{}, ids); err != nil {
		helper.RestfulError(ctx, "服务器错误："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "成功", nil)
}

// EditPassword is the edit password action of Admins controller.
func (_t *Admins) EditPassword(ctx http.Context) {
	id := ctx.Request().Input("id")
	if id == "" {
		helper.RestfulError(ctx, "请求参数错误: 缺少id")
		return
	}
	data := &models.Admin{}
	if err := facades.Orm.Query().Where("id", id).First(&data); err != nil {
		helper.RestfulError(ctx, "服务器错误："+err.Error())
		return
	}
	if data.ID == 0 {
		helper.RestfulError(ctx, "请求参数错误: 该数据不存在")
		return
	}
	var params = requests.AdminAdminEditPasswordRequest{}
	if err := ctx.Request().Bind(&params); err != nil {
		helper.RestfulError(ctx, "请求参数错误："+err.Error())
		return
	}
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "请求参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "请求参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	// Update your model here
	data.Password = helper.PasswordEncrypt(params.Password)
	if err = facades.Orm.Query().Save(&data); err != nil {
		helper.RestfulError(ctx, "服务器错误："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "更新成功", nil)
}

// Restore is the restore action of Admins controller.
//func (_t *Admins) Restore(ctx http.Context) {
//	ids := ctx.Request().QueryArray("ids")
//	if len(ids) == 0 {
//		helper.RestfulError(ctx, "请求参数错误: 缺少ids")
//		return
//	}
//	if err := facades.Orm.Query().Restore(&models.Admin{}, ids); err != nil {
//		helper.RestfulError(ctx, "服务器错误："+err.Error())
//		return
//	}
//	helper.RestfulSuccess(ctx, "请求成功", nil)
//}
