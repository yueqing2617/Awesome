package admin

import (
	"Awesome/app/http/helper"
	"Awesome/app/http/requests"
	"Awesome/app/models"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type Common struct {
	//Dependent services
}

func NewCommon() *Common {
	return &Common{
		//Inject services
	}
}

// Login is the login action of Common controller.
func (r *Common) Login(ctx http.Context) {
	var params requests.AdminLoginRequest
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
	var result struct {
		ID        uint
		Phone     string
		Nickname  string
		Avatar    string
		RoleName  string
		Role      interface{}
		Gender    string
		Remark    string
		CreatedAt string
		UpdatedAt string
	}
	_ = facades.Orm.Query().Model(&models.Admin{}).Where("phone = ? and password = ?", params.Phone, params.Password).With("Role").First(&result)
	if result.ID == 0 {
		helper.RestfulError(ctx, "账号或密码错误")
		return
	}
	// 生成token
	token, err := facades.Auth.LoginUsingID(ctx, result.ID)
	if err != nil {
		helper.RestfulError(ctx, "登录失败")
		return
	}
	helper.RestfulSuccess(ctx, "登录成功", http.Json{
		"Token": token,
		"Admin": result,
	})
}

// Logout is the logout action of Common controller.
func (r *Common) Logout(ctx http.Context) {
	// 退出登录
	_ = facades.Auth.Logout(ctx)
	helper.RestfulSuccess(ctx, "退出成功", nil)
}

// Refresh is the refresh action of Common controller.
func (r *Common) Refresh(ctx http.Context) {
	token := ctx.Request().Header("token", "")
	if token == "" {
		helper.RestfulError(ctx, "token不能为空")
		return
	}
	_ = facades.Auth.Parse(ctx, token)
	// 刷新token
	token, err := facades.Auth.Refresh(ctx)
	if err != nil {
		helper.RestfulError(ctx, "刷新失败")
		return
	}
	helper.RestfulSuccess(ctx, "刷新成功", http.Json{
		"Token": token,
	})
}

// Info is the info action of Common controller.
func (r *Common) Info(ctx http.Context) {
	token := ctx.Request().Header("token", "")
	if token == "" {
		helper.RestfulError(ctx, "token不能为空")
		return
	}
	_ = facades.Auth.Parse(ctx, token)
	// 获取当前登录用户
	var result struct {
		ID        uint
		Phone     string
		Nickname  string
		Avatar    string
		RoleName  string
		Role      interface{}
		Gender    string
		Remark    string
		CreatedAt string
		UpdatedAt string
	}
	_ = facades.Auth.User(ctx, &result)
	helper.RestfulSuccess(ctx, "获取成功", http.Json{
		"Admin": result,
	})
}

// Captcha is the captcha action of Common controller.
func (r *Common) Captcha(ctx http.Context) {
	appCode, b64s, err := helper.MakeCaptcha()
	if err != nil {
		helper.RestfulError(ctx, "验证码生成失败")
		return
	}
	helper.RestfulSuccess(ctx, "获取成功", http.Json{
		"AppCode": appCode,
		"Base64":  b64s,
	})
}
