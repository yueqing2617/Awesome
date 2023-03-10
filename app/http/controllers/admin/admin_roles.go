package admin

import (
	"Awesome/app/http/helper"
	"Awesome/app/http/requests"
	"Awesome/app/models"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type AdminRoles struct {
	//Dependent services
}

func NewAdminRoles() *AdminRoles {
	return &AdminRoles{
		//Inject services
	}
}

// Index is the index action of AdminRoles controller.
func (r *AdminRoles) Index(ctx http.Context) {
}

// Show is the show action of AdminRoles controller.
func (r *AdminRoles) Show(ctx http.Context) {
}

// Create is the creation action of AdminRoles controller.
func (r *AdminRoles) Create(ctx http.Context) {
	var params requests.AdminRoleCreateRequest
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

	// TODO: 保存数据
	data := &models.AdminRole{
		Name:        params.Name,
		DisplayName: params.DisplayName,
		Remark:      params.Remark,
		//Permissions: params.Permissions,
	}
	err = facades.Orm.Query().Create(&data)
	if err != nil {
		helper.RestfulError(ctx, "保存失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "保存成功", nil)
}

// Store is the store action of AdminRoles controller.
func (r *AdminRoles) Store(ctx http.Context) {
}
