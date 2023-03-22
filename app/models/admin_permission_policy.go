package models

import (
	"github.com/goravel/framework/database/orm"
)

type AdminPermissionPolicy struct {
	orm.Model
	Name string `gorm:"column:name;type:varchar(255);not null;default:'';comment:'名称'" json:"name"`
	Path string `gorm:"column:path;type:varchar(255);not null;default:'';comment:'路径'" json:"path"`
}

func (AdminPermissionPolicy) TableName() string {
	return "admin_permission_policies"
}
