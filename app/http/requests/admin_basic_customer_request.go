package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AdminBasicCustomerRequest struct {
	Name    string `form:"name" json:"name"`
	Code    string `form:"code" json:"code"`
	Remark  string `form:"remark" json:"remark"`
	Phone   string `form:"phone" json:"phone"`
	Company string `form:"company" json:"company"`
	Address string `form:"address" json:"address"`
	Gender  string `form:"gender" json:"gender"`
}

func (r *AdminBasicCustomerRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminBasicCustomerRequest) Rules() map[string]string {
	return map[string]string{
		"name":    "required",
		"remark":  "max_length:255",
		"phone":   "required|regex_rule:^1[3-9]\\d{9}$",
		"company": "max_length:255",
		"address": "max_length:255",
		"gender":  "in:男,女,未知",
	}
}

func (r *AdminBasicCustomerRequest) Messages() map[string]string {
	return map[string]string{
		"name.required":      "客户名称不能为空",
		"remark.max_length":  "备注不能超过255个字符",
		"phone.required":     "手机号不能为空",
		"phone.regex_rule":   "手机号格式不正确",
		"company.max_length": "公司名称不能超过255个字符",
		"address.max_length": "地址不能超过255个字符",
		"gender.in":          "性别只能是男、女或未知",
	}
}

func (r *AdminBasicCustomerRequest) Attributes() map[string]string {
	return map[string]string{
		"name":    "客户名称",
		"remark":  "备注",
		"phone":   "手机号",
		"company": "公司名称",
		"address": "地址",
		"gender":  "性别",
	}
}

func (r *AdminBasicCustomerRequest) PrepareForValidation(data validation.Data) error {
	return nil
}
