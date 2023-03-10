package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AdminAdminEditPasswordRequest struct {
	Password string `form:"password" json:"password"`
	Confirm  string `form:"confirm" json:"confirm"`
}

func (r *AdminAdminEditPasswordRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminAdminEditPasswordRequest) Rules() map[string]string {
	return map[string]string{
		"password": "required|min_len:6|max_length:255",
		"confirm":  "required|eq_field:password",
	}
}

func (r *AdminAdminEditPasswordRequest) Messages() map[string]string {
	return map[string]string{
		"password.required": "密码不能为空",
		"password.min_len":  "密码不能少于6个字符",
		"password.max_len":  "密码不能超过255个字符",
		"confirm.required":  "确认密码不能为空",
		"confirm.eq_field":  "两次密码不一致",
	}
}

func (r *AdminAdminEditPasswordRequest) Attributes() map[string]string {
	return map[string]string{
		"password": "密码",
		"confirm":  "确认密码",
	}
}

func (r *AdminAdminEditPasswordRequest) PrepareForValidation(data validation.Data) error {
	return nil
}
