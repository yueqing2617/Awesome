package requests

import (
	"errors"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AdminBasicCustomerRequest struct {
	Name    string `form:"name" json:"name"`
	Code    string `form:"code" json:"code"`
	Phone   string `form:"phone" json:"phone"`
	Company string `form:"company" json:"company"`
	Address string `form:"address" json:"address"`
	Gender  string `form:"gender" json:"gender"`
	Remark  string `form:"remark" json:"remark"`
}

func (r *AdminBasicCustomerRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminBasicCustomerRequest) Rules(ctx http.Context) map[string]string {
	method := ctx.Request().Method()
	switch method {
	case "POST":
		return map[string]string{
			"name":    "required|unique:basic_customers,name",
			"code":    "max_length:255",
			"phone":   "phone|unique:basic_customers,phone",
			"company": "max_length:255",
			"address": "max_length:255",
			"gender":  "in:男,女,未知",
			"remark":  "max_length:255",
		}
	case "PUT":
		return map[string]string{
			"name":    "required|not_exists:basic_customers,name",
			"code":    "max_length:255",
			"phone":   "phone",
			"company": "max_length:255",
			"address": "max_length:255",
			"gender":  "in:男,女,未知",
			"remark":  "max_length:255",
		}
	default:
		return map[string]string{}
	}
}

func (r *AdminBasicCustomerRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"name.required":      "名称不能为空",
		"name.unique":        "名称已存在",
		"name.not_exists":    "名称不能更改",
		"code.max_length":    "编码不能超过255个字符",
		"phone.regex_rule":   "电话格式不正确",
		"company.max_length": "公司不能超过255个字符",
		"address.max_length": "地址不能超过255个字符",
		"genger.in":          "性别只能是男、女或未知",
		"remark.max_length":  "备注不能超过255个字符",
	}
}

func (r *AdminBasicCustomerRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"name":    "名称",
		"code":    "编码",
		"phone":   "电话",
		"company": "公司",
		"address": "地址",
		"gender":  "性别",
		"remark":  "备注",
	}
}

func (r *AdminBasicCustomerRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
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
