package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AdminBasicRequest struct {
	Name   string `form:"name" json:"name"`
	Code   string `form:"code" json:"code"`
	Remark string `form:"remark" json:"remark"`
}

func (r *AdminBasicRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminBasicRequest) Rules() map[string]string {
	return map[string]string{
		"name":   "required",
		"code":   "max_length:10",
		"remark": "max_length:255",
	}
}

func (r *AdminBasicRequest) Messages() map[string]string {
	return map[string]string{
		"name.required":     "名称不能为空",
		"code.max_length":   "编码长度不能超过10",
		"remark.max_length": "备注长度不能超过255",
	}
}

func (r *AdminBasicRequest) Attributes() map[string]string {
	return map[string]string{
		"name":   "名称",
		"code":   "编码",
		"remark": "备注",
	}
}

func (r *AdminBasicRequest) PrepareForValidation(data validation.Data) error {
	return nil
}
