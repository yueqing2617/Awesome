package requests

import (
	"errors"
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

func (r *AdminBasicRequest) Rules(ctx http.Context) map[string]string {
	method := ctx.Request().Method()
	switch method {
	case "POST":
		return map[string]string{
			"name":   "required",
			"code":   "max_length:255",
			"remark": "max_length:255",
		}
	case "PUT":
		return map[string]string{
			"name":   "required",
			"code":   "max_length:255",
			"remark": "max_length:255",
		}
	default:
		return map[string]string{
			"path": "required",
		}
	}
}

func (r *AdminBasicRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"name.required":     "名称不能为空",
		"code.max_length":   "编码长度不能超过255",
		"remark.max_length": "备注长度不能超过255",
	}
}

func (r *AdminBasicRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"name":   "名称",
		"code":   "编码",
		"remark": "备注",
	}
}

func (r *AdminBasicRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
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
