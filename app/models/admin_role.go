package models

import (
	"github.com/goravel/framework/database/orm"
	"strings"
)

type StringArray []string

type AdminRole struct {
	orm.Model
	Name          string   `gorm:"column:name;type:varchar(255);not null;default:'';comment:'名称'" json:"name"`
	DisplayName   string   `gorm:"column:display_name;type:varchar(255);not null;default:'';comment:'显示名称'" json:"display_name"`
	Remark        string   `gorm:"column:remark;type:varchar(255);default:'';comment:'备注'" json:"remark"`
	PermissionStr string   `gorm:"column:permissions;type:text;default:'';comment:'权限'" json:"-"`
	Permissions   []string `gorm:"-" json:"permissions"`
	orm.SoftDeletes
}

func (AdminRole) TableName() string {
	return "admin_roles"
}

// 处理取出的permissionStr
func (r *AdminRole) Scan() error {
	if r.PermissionStr != "" {
		// 从数据库中取出的是字符串，需要转换成[]string
		var permissions []string
		// string转[]string,以逗号分隔并且去掉空格
		permissions = strings.Split(strings.Replace(r.PermissionStr, " ", "", -1), ",")
		// 过滤空字符串
		for i := 0; i < len(permissions); i++ {
			if permissions[i] == "" {
				permissions = append(permissions[:i], permissions[i+1:]...)
			}
		}
		r.Permissions = permissions
	}
	return nil
}
