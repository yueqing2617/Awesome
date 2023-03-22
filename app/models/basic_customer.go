package models

import (
	"github.com/goravel/framework/database/orm"
)

type BasicCustomer struct {
	orm.Model
	Name    string `gorm:"column:name;type:varchar(255);not null;default:'';comment:'名称'" json:"name"`
	Code    string `gorm:"column:code;type:varchar(255);default:'';comment:'编码'" json:"code"`
	Phone   string `gorm:"column:phone;type:varchar(11);default:'';comment:'电话'" json:"phone"`
	Company string `gorm:"column:company;type:varchar(255);default:'';comment:'公司'" json:"company"`
	Address string `gorm:"column:address;type:varchar(255);default:'';comment:'地址'" json:"address"`
	Gender  string `gorm:"column:gender;type:varchar(5);default:'';comment:'性别'" json:"gender"`
	Remark  string `gorm:"column:remark;type:varchar(255);default:'';comment:'备注'" json:"remark"`
}

func (BasicCustomer) TableName() string {
	return "basic_customers"
}
