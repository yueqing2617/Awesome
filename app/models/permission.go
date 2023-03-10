package models

import (
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
)

type Permission struct {
	orm.Model
	Name     string       `gorm:"type:varchar(255);not null"`
	Path     string       `gorm:"type:varchar(255);not null;unique"`
	Pid      uint         `gorm:"type:int(11);not null;default:0"`
	Children []Permission `gorm:"-"`
}

// 无线递归,一次性获取所有数据
func (p *Permission) GetPermissionList() ([]Permission, error) {
	var permissions []Permission
	err := facades.Orm.Query().Order("sort asc").Find(&permissions)
	if err != nil {
		return nil, err
	}
	// 递归
	data := listToTree(permissions, 0)
	return data, nil
}

// listToTree 将权限列表转换为树形结构
// parentId: 父级ID
func listToTree(permissions []Permission, parentId uint) []Permission {
	var tree []Permission
	for _, permission := range permissions {
		if permission.Pid == parentId {
			permission.Children = listToTree(permissions, permission.ID)
			tree = append(tree, permission)
		}
	}
	return tree
}
