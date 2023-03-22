package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AdminLoginRequest struct {
	Phone     string `form:"phone" json:"phone"`
	Password  string `form:"password" json:"password"`
	AppSecret string `form:"app_secret" json:"app_secret"`
	Captcha   string `form:"captcha" json:"captcha"`
}

func (r *AdminLoginRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminLoginRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"phone":      "required|phone",
		"password":   "required|min_len:6",
		"app_secret": "required",
		"captcha":    "required|captcha",
	}
}

func (r *AdminLoginRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"phone.required":      "手机号不能为空",
		"phone.phone":         "手机号格式不正确",
		"password.required":   "密码不能为空",
		"password.min_len":    "密码不能少于6个字符",
		"app_secret.required": "app_secret不能为空",
		"captcha.required":    "验证码不能为空",
	}
}

func (r *AdminLoginRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"phone":      "手机号",
		"password":   "密码",
		"app_secret": "app_secret",
		"captcha":    "验证码",
	}
}

func (r *AdminLoginRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	//pwd, _ := data.Get("password")
	//if pwd != nil {
	//	hash, err := facades.Hash.Make(pwd.(string))
	//	if err != nil {
	//		return err
	//	}
	//	_ = data.Set("password", hash)
	//}
	return nil
}
