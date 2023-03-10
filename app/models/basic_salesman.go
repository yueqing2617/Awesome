package models

import (
	"github.com/goravel/framework/database/orm"
)

type BasicSalesman struct {
	orm.Model
	Name   string `gorm:"type:varchar(255);not null;unique"`
	Code   string `gorm:"type:varchar(255);"`
	Remark string `gorm:"type:varchar(255);"`
}
