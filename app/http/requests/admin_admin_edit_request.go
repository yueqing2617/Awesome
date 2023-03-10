package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AdminAdminEditRequest struct {
	Email    string `form:"email" json:"email"`
	Nickname string `form:"nickname" json:"nickname"`
	Gender   string `form:"gender" json:"gender"`
	Avatar   string `form:"avatar" json:"avatar"`
	RoleName string `form:"role_name" json:"role_name"`
	Remark   string `form:"remark" json:"remark"`
}

func (r *AdminAdminEditRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminAdminEditRequest) Rules() map[string]string {
	return map[string]string{
		"email":     "required|email",
		"nickname":  "required|max_length:255",
		"gender":    "required|in:男,女,保密",
		"avatar":    "url",
		"role_name": "required",
		"remark":    "max_length:255",
	}
}

func (r *AdminAdminEditRequest) Messages() map[string]string {
	return map[string]string{
		"email.required":      "邮箱不能为空",
		"email.email":         "邮箱格式不正确",
		"nickname.required":   "昵称不能为空",
		"nickname.max_length": "昵称长度不能超过255",
		"gender.required":     "性别不能为空",
		"gender.in":           "性别不正确",
		"avatar.url":          "头像链接不正确",
		"role_name.required":  "角色不能为空",
		"remark.max_length":   "备注长度不能超过255",
	}
}

func (r *AdminAdminEditRequest) Attributes() map[string]string {
	return map[string]string{
		"email":    "邮箱",
		"nickname": "昵称",
		"gender":   "性别",
		"avatar":   "头像",
		"role_id":  "角色",
		"remark":   "备注",
	}
}

func (r *AdminAdminEditRequest) PrepareForValidation(data validation.Data) error {
	return nil
}
