package models

import (
	"encoding/json"
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
)

type AdminPermission struct {
	orm.Model
	Name      string `gorm:"column:name;type:varchar(255);not null;default:'';comment:'名称'"`
	Path      string `gorm:"column:path;type:varchar(255);not null;default:'';comment:'路径'"`
	ParentId  uint   `gorm:"column:parent_id;type:int(11);not null;default:0;comment:'父级ID'"`
	Remark    string `gorm:"column:remark;type:varchar(255);default:'';comment:'备注'"`
	Meta      *Meta  `gorm:"-"`
	MetaJSON  string `gorm:"column:meta;type:json;default:'';comment:'元数据'"`
	Redirect  string `gorm:"column:redirect;type:varchar(255);default:'';comment:'重定向'"`
	Active    string `gorm:"column:active;type:varchar(255);default:'';comment:'高亮'"`
	Component string `gorm:"column:component;type:varchar(255);default:'';comment:'组件'"`
	Sort      uint   `gorm:"column:sort;type:int(11);not null;default:0;comment:'排序'"`
}

type Meta struct {
	Icon             string `json:"icon" form:"icon"`
	Title            string `json:"title" form:"title"`
	Type             string `json:"type" form:"type"`
	Color            string `json:"color" form:"color"`
	Hidden           bool   `json:"hidden" form:"hidden"`
	HiddenBreadcrumb bool   `json:"hiddenBreadcrumb" form:"meta.hiddenBreadcrumb"`
	FullPage         bool   `json:"fullpage" form:"fullpage"`
	Tag              string `json:"tag" form:"tag"`
	Affix            bool   `json:"affix" form:"affix"`
}

type PermissionResult struct {
	ID         uint               `json:"id"`
	Name       string             `json:"name"`
	Path       string             `json:"path"`
	ParentId   uint               `json:"parent_id"`
	Remark     string             `json:"remark"`
	Meta       *Meta              `json:"meta"`
	Redirect   string             `json:"redirect"`
	Component  string             `json:"component"`
	Active     string             `json:"active"`
	Sort       uint               `json:"sort"`
	Children   []PermissionResult `json:"children"`
	CreateTime int64              `json:"create_time"`
	UpdateTime int64              `json:"update_time"`
}

func (AdminPermission) TableName() string {
	return "admin_permissions"
}

func (p *AdminPermission) GetPermissionList(pid uint) ([]PermissionResult, error) {
	var permissions []AdminPermission
	err := facades.Orm.Query().Model(&p).Order("sort asc").Find(&permissions)
	if err != nil {
		return nil, err
	}
	var data []PermissionResult
	for _, permission := range permissions {
		var meta *Meta
		if permission.MetaJSON != "" {
			meta = &Meta{}
			err := json.Unmarshal([]byte(permission.MetaJSON), meta)
			if err != nil {
				return nil, err
			}
		}
		data = append(data, PermissionResult{
			ID:         permission.ID,
			Name:       permission.Name,
			Path:       permission.Path,
			ParentId:   permission.ParentId,
			Remark:     permission.Remark,
			Meta:       meta,
			Redirect:   permission.Redirect,
			Component:  permission.Component,
			Sort:       permission.Sort,
			CreateTime: permission.CreatedAt,
			UpdateTime: permission.UpdatedAt,
		})
	}
	//data := PermissionsToTree(permissions, pid)
	data = PermissionsToTree(data, pid)
	return data, nil
}

// PermissionsToTree 将权限列表转换为树形结构
func PermissionsToTree(permissions []PermissionResult, parentId uint) []PermissionResult {
	var tree []PermissionResult
	for _, permission := range permissions {
		if permission.ParentId == parentId {
			permission.Children = PermissionsToTree(permissions, permission.ID)
			tree = append(tree, permission)
		}
	}
	return tree
}
