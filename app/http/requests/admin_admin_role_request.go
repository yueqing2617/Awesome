package requests

import (
	"errors"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AdminAdminRoleRequest struct {
	Name        string   `form:"name" json:"name"`
	DisplayName string   `form:"display_name" json:"display_name"`
	Remark      string   `form:"remark" json:"remark"`
	Permissions []string `form:"permissions" json:"permissions"`
}

func (r *AdminAdminRoleRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminAdminRoleRequest) Rules(ctx http.Context) map[string]string {
	method := ctx.Request().Method()
	switch method {
	case "POST":
		return map[string]string{
			"name":         "required|min_len:2|max_len:255|alpha|unique:admin_roles,name",
			"display_name": "required|min_len:2|max_len:255",
			"remark":       "max_length:255",
			"permissions":  "required|min_len:1",
		}
	case "PUT":
		return map[string]string{
			"name":         "required|min_len:2|max_len:255|alpha",
			"display_name": "required|min_len:2|max_len:255",
			"remark":       "max_length:255",
			"permissions":  "required|min_len:1",
		}
	default:
		return map[string]string{}
	}
}

func (r *AdminAdminRoleRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"name.required":         "角色名称不能为空",
		"name.min_len":          "角色名称不能少于2个字符",
		"name.max_len":          "角色名称不能多于255个字符",
		"name.alpha":            "角色名称只能包含字母",
		"name.unique":           "角色名称已存在",
		"display_name.required": "显示名称不能为空",
		"display_name.min_len":  "显示名称不能少于2个字符",
		"display_name.max_len":  "显示名称不能多于255个字符",
		"remark.max_length":     "备注不能多于255个字符",
		"permissions.required":  "权限不能为空",
		"permissions.min_len":   "权限不能少于1个",
	}
}

func (r *AdminAdminRoleRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"name":         "角色名称",
		"display_name": "显示名称",
		"remark":       "备注",
		"permissions":  "权限",
	}
}

func (r *AdminAdminRoleRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
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
