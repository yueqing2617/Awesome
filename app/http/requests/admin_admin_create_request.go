package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AdminAdminCreateRequest struct {
	Phone           string `form:"phone" json:"phone"`
	Password        string `form:"password" json:"password"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password"`
	Email           string `form:"email" json:"email"`
	Nickname        string `form:"nickname" json:"nickname"`
	Gender          string `form:"gender" json:"gender"`
	Avatar          string `form:"avatar" json:"avatar"`
	RoleName        string `form:"role_name" json:"role_name"`
	Remark          string `form:"remark" json:"remark"`
}

func (r *AdminAdminCreateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminAdminCreateRequest) Rules() map[string]string {
	return map[string]string{
		"phone":            "required|regex_rule:^1[3-9]\\d{9}$|unique:admins,phone",
		"password":         "required|min_len:6",
		"confirm_password": "required|eq_field:password",
		"email":            "required|email",
		"nickname":         "required",
		"gender":           "required|in:男,女,保密",
		"avatar":           "required|url",
		"role_name":        "required",
		"remark":           "max_length:2",
	}
}

func (r *AdminAdminCreateRequest) Messages() map[string]string {
	return map[string]string{
		"phone.required":            "手机号不能为空",
		"phone.regex_rule":          "手机号格式不正确",
		"phone.unique":              "手机号已存在",
		"password.required":         "密码不能为空",
		"password.min_len":          "密码不能少于6个字符",
		"confirm_password.required": "确认密码不能为空",
		"confirm_password.eq":       "两次密码不一致",
		"email.required":            "邮箱不能为空",
		"email.email":               "邮箱格式不正确",
		"nickname.required":         "昵称不能为空",
		"gender.required":           "性别不能为空",
		"gender.in":                 "请选择正确的性别",
		"avatar.required":           "头像不能为空",
		"avatar.url":                "头像地址不正确",
		"role_name.required":        "角色不能为空",
		"remark.max_length":         "备注不能超过255个字符",
	}
}

func (r *AdminAdminCreateRequest) Attributes() map[string]string {
	return map[string]string{
		"phone":            "手机号",
		"password":         "密码",
		"confirm_password": "确认密码",
		"email":            "邮箱",
		"nickname":         "昵称",
		"gender":           "性别",
		"avatar":           "头像",
		"role_id":          "角色",
		"remark":           "备注",
	}
}

func (r *AdminAdminCreateRequest) PrepareForValidation(data validation.Data) error {
	if gender, exist := data.Get("gender"); exist {
		return data.Set("gender", gender)
	}
	return nil
}
