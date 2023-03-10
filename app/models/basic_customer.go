package models

import (
	"github.com/goravel/framework/database/orm"
)

type BasicCustomer struct {
	orm.Model
	Name    string `gorm:"type:varchar(255);not null;unique"`
	Code    string `gorm:"type:varchar(255);"`
	Gender  string `gorm:"type:varchar(4);"`
	Phone   string `gorm:"type:varchar(11);"`
	Company string `gorm:"type:varchar(255);"`
	Address string `gorm:"type:varchar(255);"`
	Remark  string `gorm:"type:varchar(255);"`
}
