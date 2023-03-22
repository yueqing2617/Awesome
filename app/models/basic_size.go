package models

import (
	"github.com/goravel/framework/database/orm"
)

type BasicSize struct {
	orm.Model
	Name   string `gorm:"column:name;type:varchar(255);not null;default:'';comment:'名称'" json:"name"`
	Code   string `gorm:"column:code;type:varchar(255);default:'';comment:'编码'" json:"code"`
	Remark string `gorm:"column:remark;type:varchar(255);default:'';comment:'备注'" json:"remark"`
}

func (BasicSize) TableName() string {
	return "basic_sizes"
}
