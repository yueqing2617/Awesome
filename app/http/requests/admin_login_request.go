package requests

import (
	"Awesome/app/http/helper"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AdminLoginRequest struct {
	Phone    string `form:"phone" json:"phone"`
	Password string `form:"password" json:"password"`
	AppCode  string `form:"app_code" json:"app_code"`
	Captcha  string `form:"captcha" json:"captcha"`
}

func (r *AdminLoginRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminLoginRequest) Rules() map[string]string {
	return map[string]string{
		"phone":    "required|regex_rule:^1[3456789]\\d{9}$",
		"password": "required",
		"app_code": "required",
		"captcha":  "required|captcha",
	}
}

func (r *AdminLoginRequest) Messages() map[string]string {
	return map[string]string{
		"phone.required":    "手机号不能为空",
		"phone.regex_rule":  "手机号格式不正确",
		"password.required": "密码不能为空",
		"app_code.required": "安全码不能为空",
		"captcha.required":  "验证码不能为空",
		"captcha.captcha":   "验证码错误",
	}
}

func (r *AdminLoginRequest) Attributes() map[string]string {
	return map[string]string{
		"phone":    "手机号",
		"password": "密码",
		"app_code": "安全码",
		"captcha":  "验证码",
	}
}

func (r *AdminLoginRequest) PrepareForValidation(data validation.Data) error {
	pwd, _ := data.Get("password")
	if pwd != nil {
		_ = data.Set("password", helper.PasswordEncrypt(pwd.(string)))
	}
	return nil
}
