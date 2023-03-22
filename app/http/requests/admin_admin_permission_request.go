package requests

import (
	"Awesome/app/models"
	"errors"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AdminAdminPermissionRequest struct {
	Name      string      `form:"name" json:"name"`
	Path      string      `form:"path" json:"path"`
	ParentId  uint        `form:"parentId" json:"parentId"`
	Remark    string      `form:"remark" json:"remark"`
	Meta      models.Meta `form:"meta" json:"meta"`
	Redirect  string      `form:"redirect" json:"redirect"`
	Component string      `form:"component" json:"component"`
	Active    string      `form:"active" json:"active"`
}

func (r *AdminAdminPermissionRequest) Authorize(ctx http.Context) error {
	method := ctx.Request().Method()
	switch method {
	case "POST":
		return nil
	case "PUT":
		return nil
	case "DELETE":
		return nil
	default:
		return nil
	}
}

func (r *AdminAdminPermissionRequest) Rules(ctx http.Context) map[string]string {
	method := ctx.Request().Method()
	switch method {
	case "POST":
		return map[string]string{
			"name":                  "required|min_len:2|max_len:255|unique:admin_permissions,name",
			"path":                  "required|max_len:255|unique:admin_permissions,path",
			"parentId":              "numeric",
			"remark":                "max_len:255",
			"meta.title":            "required|min_len:2|max_len:255",
			"meta.icon":             "max_len:255",
			"meta.color":            "max_len:255",
			"redirect":              "max_len:255",
			"component":             "max_len:255",
			"active":                "max_len:255",
			"meta.type":             "required|in:menu,button,iframe,link",
			"meta.tag":              "max_len:255",
			"meta.hidden":           "bool",
			"meta.fullpage":         "bool",
			"meta.hiddenBreadcrumb": "bool",
		}
	case "PUT":
		return map[string]string{
			"name":                  "required|min_len:2|max_len:255",
			"path":                  "required|max_len:255|not_exists:admin_permissions,path",
			"parentId":              "numeric",
			"remark":                "max_len:255",
			"meta.title":            "required|min_len:2|max_len:255",
			"meta.icon":             "max_len:255",
			"meta.color":            "max_len:255",
			"redirect":              "max_len:255",
			"component":             "max_len:255",
			"active":                "max_len:255",
			"meta.type":             "required|in:menu,button,iframe,link",
			"meta.tag":              "max_len:255",
			"meta.hidden":           "bool",
			"meta.fullpage":         "bool",
			"meta.hiddenBreadcrumb": "bool",
		}
	default:
		return nil
	}
}

func (r *AdminAdminPermissionRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"name.required":              "别名不能为空",
		"name.min_len":               "别名长度不能小于2",
		"name.max_len":               "别名长度不能大于255",
		"name.unique":                "别名已存在",
		"path.required":              "路由地址不能为空",
		"path.max_len":               "路由地址长度不能大于255",
		"path.unique":                "路由地址已存在",
		"path.not_exists":            "路由地址不能更改",
		"parentId.numeric":           "上级菜单ID必须为数字",
		"remark.max_len":             "备注长度不能大于255",
		"meta.title.required":        "菜单标题不能为空",
		"meta.title.min_len":         "菜单标题长度不能小于2",
		"meta.title.max_len":         "菜单标题长度不能大于255",
		"meta.icon.max_len":          "菜单图标长度不能大于255",
		"meta.color.max_len":         "菜单颜色长度不能大于255",
		"redirect.max_len":           "重定向地址长度不能大于255",
		"component.max_len":          "组件地址长度不能大于255",
		"active.max_len":             "高亮菜单长度不能大于255",
		"meta.type.required":         "菜单类型不能为空",
		"meta.type.in":               "菜单类型不正确",
		"meta.tag.max_len":           "菜单标签长度不能大于255",
		"meta.hidden.bool":           "是否隐藏带单必须为布尔值",
		"meta.fullpage.bool":         "整页显示必须为布尔值",
		"meta.hiddenBreadcrumb.bool": "是否隐藏面包屑必须为布尔值",
	}
}

func (r *AdminAdminPermissionRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"name":                  "别名",
		"path":                  "路由地址",
		"parentId":              "上级菜单ID",
		"remark":                "备注",
		"meta":                  "菜单元数据",
		"redirect":              "重定向地址",
		"component":             "组件地址",
		"active":                "高亮菜单",
		"meta.title":            "菜单标题",
		"meta.icon":             "菜单图标",
		"meta.color":            "菜单颜色",
		"meta.hidden":           "是否隐藏带单",
		"meta.hiddenBreadcrumb": "是否隐藏面包屑",
		"meta.fullpage":         "整页路由",
		"meta.tag":              "标签",
		"meta.type":             "类型",
	}
}

func (r *AdminAdminPermissionRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	method := ctx.Request().Method()
	switch method {
	case "POST":
		opt, _ := data.Get("type")
		if opt == "menu" {
			_ = data.Set("redirect", "")
			component, _ := data.Get("component")
			if component == "" {
				return errors.New("权限类型为menu时，组件地址不能为空,例如： home/widgets")
			}
		} else if opt == "button" {
			path, _ := data.Get("path")
			if path == "" {
				return errors.New("权限类型为按钮时，路由地址不能为空")
			}
			_ = data.Set("redirect", "")
			_ = data.Set("component", "")
		} else if opt == "iframe" {
			component, _ := data.Get("component")
			if component == "" {
				return errors.New("权限类型为iframe时，组件地址不能为空")
			}
			_ = data.Set("redirect", "")
		} else if opt == "link" {
			redirect, _ := data.Get("redirect")
			if redirect == "" {
				return errors.New("权限类型为link时，重定向地址不能为空")
			}
			_ = data.Set("component", "")
		}
		return nil
	case "PUT":
		id := ctx.Request().Input("id")
		if id == "" {
			return errors.New("id不能为空")
		}
		opt, _ := data.Get("type")
		if opt == "menu" {
			_ = data.Set("redirect", "")
		} else if opt == "button" {
			path, _ := data.Get("path")
			if path == "" {
				return errors.New("权限类型为按钮时，路由地址不能为空")
			}
			_ = data.Set("redirect", "")
			_ = data.Set("component", "")
		} else if opt == "iframe" {
			component, _ := data.Get("component")
			if component == "" {
				return errors.New("权限类型为iframe时，组件地址不能为空")
			}
			_ = data.Set("redirect", "")
		} else if opt == "link" {
			redirect, _ := data.Get("redirect")
			if redirect == "" {
				return errors.New("权限类型为link时，重定向地址不能为空")
			}
			_ = data.Set("component", "")
		}
		return nil
	default:
		return nil
	}
}
