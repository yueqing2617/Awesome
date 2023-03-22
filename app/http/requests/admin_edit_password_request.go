package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AdminEditPasswordRequest struct {
	OldPassword     string `json:"old_password" form:"old_password"`
	NewPassword     string `json:"new_password" form:"new_password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

func (r *AdminEditPasswordRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminEditPasswordRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"old_password":     "required",
		"new_password":     "required|min_len:6|max_len:36",
		"confirm_password": "required|eq_field:new_password",
	}
}

func (r *AdminEditPasswordRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"old_password.required":     "旧密码不能为空",
		"new_password.required":     "新密码不能为空",
		"new_password.min_len":      "新密码长度不能小于6",
		"new_password.max_len":      "新密码长度不能大于36",
		"confirm_password.required": "确认密码不能为空",
		"confirm_password.eq_field": "两次密码不一致",
	}
}

func (r *AdminEditPasswordRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"old_password":     "旧密码",
		"new_password":     "新密码",
		"confirm_password": "确认密码",
	}
}

func (r *AdminEditPasswordRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	method := ctx.Request().Method()
	switch method {
	case "POST":
		return nil
	case "PUT":
		return nil
	default:
		return nil
	}
}
