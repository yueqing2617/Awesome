package models

import (
	"github.com/goravel/framework/database/orm"
)

type Admin struct {
	orm.Model
	Nickname string     `gorm:"column:nickname;type:varchar(255);default:'';comment:'昵称'" json:"nickname"`
	Phone    string     `gorm:"column:phone;type:varchar(11);default:'';comment:'手机号'" json:"phone"`
	Password string     `gorm:"column:password;type:varchar(255);default:'';comment:'密码'" json:"password"`
	Avatar   string     `gorm:"column:avatar;type:text;default:'';comment:'头像'" json:"avatar"`
	Email    string     `gorm:"column:email;type:varchar(255);default:'';comment:'邮箱'" json:"email"`
	Remark   string     `gorm:"column:remark;type:varchar(255);default:'';comment:'备注'" json:"remark"`
	RoleID   uint       `gorm:"column:role_id;type:int(11);default:0;comment:'角色ID'" json:"role_id"`
	Role     *AdminRole `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	Gender   string     `gorm:"column:gender;type:varchar(5);default:'未知';comment:'性别'" json:"gender"`
	orm.SoftDeletes
}

func (Admin) TableName() string {
	return "admins"
}
