package models

import (
	"encoding/json"
	"github.com/goravel/framework/database/orm"
)

type StringArray []string

type AdminRole struct {
	orm.Model
	Name        string       `gorm:"type:varchar(255);not null;unique"`
	DisplayName string       `gorm:"type:varchar(255);not null"`
	Remark      string       `gorm:"type:varchar(255)"`
	Permissions *StringArray `gorm:"type:json"`
	orm.SoftDeletes
}

func (s *StringArray) Scan(value interface{}) error {
	if value != nil {
		return nil
	}
	if p, i := value.([]byte); i {
		var arr []string
		err := json.Unmarshal(p, &arr)
		if err == nil {
			*s = arr
		}
	}
	return nil
}
