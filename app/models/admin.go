package models

import (
	"github.com/goravel/framework/database/orm"
)

type Admin struct {
	orm.Model
	Phone    string     `gorm:"type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(32);not null"`
	Email    string     `gorm:"type:varchar(255)"`
	Gender   string     `gorm:"type:varchar(4) DEFAULT '未知'"` // 未知、男、女
	Nickname string     `gorm:"type:varchar(255)"`
	Avatar   string     `gorm:"type:varchar(255) DEFAULT 'https://avatars.githubusercontent.com/u/16495509?v=4'"`
	RoleName string     `gorm:"type:varchar(255) DEFAULT 'guest'"`
	Role     *AdminRole `gorm:"foreignKey:RoleName;references:Name"`
	Remark   string     `gorm:"type:varchar(255)"`
	orm.SoftDeletes
}
