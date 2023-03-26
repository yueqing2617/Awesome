package models

import (
	"github.com/goravel/framework/database/orm"
)

type ClothTailor struct {
	orm.Model
	ClothOrderCode string                     `gorm:"column:cloth_order_code;type:varchar(255);not null;comment:生产订单号" json:"cloth_order_code"`
	ClothOrder     *ClothOrder                `gorm:"foreignKey:cloth_order_code;references:code;" json:"cloth_order"`
	ClothStyleCode string                     `gorm:"column:cloth_style_code;type:varchar(255);not null;comment:布料款式编码" json:"cloth_style_code"`
	Total          uint                       `gorm:"column:total;type:int(10) unsigned;not null;comment:总数" json:"total"`
	CompletedNum   uint                       `gorm:"column:completed_num;type:int(10) unsigned;not null;comment:已完成数量" json:"completed_num"`
	IsCompleted    bool                       `gorm:"column:is_completed;type:tinyint(1);not null;comment:是否完成" json:"is_completed"`
	CuttingPieces  []ClothTailorCuttingPieces `gorm:"foreignKey:cloth_tailor_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"cloth_tailor_cutting"`
	orm.SoftDeletes
}

func (ClothTailor) TableName() string {
	return "cloth_tailors"
}
