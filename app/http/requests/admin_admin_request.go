package requests

import (
	"errors"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AdminAdminRequest struct {
	Phone           string `form:"phone" json:"phone"`
	Nickname        string `form:"nickname" json:"nickname"`
	Avatar          string `form:"avatar" json:"avatar"`
	Email           string `form:"email" json:"email"`
	Remark          string `form:"remark" json:"remark"`
	RoleID          uint   `form:"roleId" json:"roleId"`
	Gender          string `json:"gender" form:"gender"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

func (r *AdminAdminRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminAdminRequest) Rules(ctx http.Context) map[string]string {
	method := ctx.Request().Method()
	switch method {
	case "POST":
		return map[string]string{
			"phone":            "required|phone",
			"nickname":         "required|min_len:2|max_len:20",
			"email":            "required|email",
			"remark":           "max_length:255",
			"roleId":           "required|numeric",
			"gender":           "required|in:男,女,未知",
			"password":         "required|min_len:6|max_len:36",
			"confirm_password": "required|eq_field:password",
		}
	case "PUT":
		return map[string]string{
			"nickname": "required|min_len:2|max_len:20",
			"email":    "required|email",
			"remark":   "max_length:255",
			"roleId":   "required|numeric",
			"gender":   "required|in:男,女,未知",
		}
	default:
		return map[string]string{}
	}
}

func (r *AdminAdminRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"phone.required":            "手机号不能为空",
		"phone.phone":               "手机号格式不正确",
		"nickname.required":         "昵称不能为空",
		"nickname.min_len":          "昵称不能少于2个字符",
		"nickname.max_len":          "昵称不能多于20个字符",
		"avatar.required":           "头像不能为空",
		"avatar.url":                "头像链接地址不正确",
		"email.required":            "邮箱不能为空",
		"email.email":               "邮箱格式不正确",
		"remark.max_length":         "备注不能多于255个字符",
		"role_id.required":          "角色ID不能为空",
		"role_id.numeric":           "角色ID必须为数字",
		"gender.required":           "性别不能为空",
		"gender.in":                 "性别只能为男、女或未知",
		"password.required":         "密码不能为空",
		"password.min_len":          "密码不能少于6个字符",
		"password.max_len":          "密码不能多于36个字符",
		"confirm_password.required": "确认密码不能为空",
		"confirm_password.eq_field": "两次密码输入不一致",
	}
}

func (r *AdminAdminRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"phone":            "手机号",
		"nickname":         "昵称",
		"avatar":           "头像",
		"email":            "邮箱",
		"remark":           "备注",
		"role_id":          "角色ID",
		"gender":           "性别",
		"password":         "密码",
		"confirm_password": "确认密码",
	}
}

func (r *AdminAdminRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	method := ctx.Request().Method()
	switch method {
	case "POST":
		return nil
	case "PUT":
		id := ctx.Request().Input("id")
		if id == "" {
			return errors.New("id不能为空")
		}
		return nil
	default:
		return nil
	}
}
