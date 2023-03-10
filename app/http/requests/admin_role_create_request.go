package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AdminRoleCreateRequest struct {
	Name        string   `form:"name" json:"name"`
	DisplayName string   `form:"display_name" json:"display_name"`
	Remark      string   `form:"remark" json:"remark"`
	Permissions []string `form:"permissions" json:"permissions"`
}

func (r *AdminRoleCreateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminRoleCreateRequest) Rules() map[string]string {
	return map[string]string{
		"name":         "required",
		"display_name": "required",
		"permissions":  "required|array|min_len:1",
	}
}

func (r *AdminRoleCreateRequest) Messages() map[string]string {
	return map[string]string{
		"name.required":         "角色名称不能为空",
		"display_name.required": "角色显示名称不能为空",
		"permissions.required":  "权限不能为空",
		"permissions.array":     "权限必须是数组",
		"permissions.min_len":   "权限至少选择一个",
	}
}

func (r *AdminRoleCreateRequest) Attributes() map[string]string {
	return map[string]string{
		"name":         "角色名称",
		"display_name": "角色显示名称",
		"permissions":  "权限",
	}
}

func (r *AdminRoleCreateRequest) PrepareForValidation(data validation.Data) error {
	return nil
}
