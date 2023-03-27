package models

import "github.com/goravel/framework/database/orm"

type ClothTailorCuttingPieceProcess struct {
	orm.Model
	Name           string  `json:"name" form:"name" gorm:"column:name;type:varchar(255);not null;comment:工序名称"`
	Sort           uint    `json:"sort" form:"sort" gorm:"column:sort;type:int(10) unsigned;not null;default:0;comment:排序"`
	Price          float64 `json:"price" form:"price" gorm:"column:price;type:decimal(10,2);not null;default:0.00;comment:单价"`
	CuttingPieceID uint    `gorm:"column:cutting_piece_id;type:int(10) unsigned;not null;comment:裁剪订单ID" json:"cutting_piece_id"`
	IsCompleted    bool    `json:"is_completed" form:"is_completed" gorm:"column:is_completed;type:tinyint(1);not null;default:0;comment:是否完成"`
	EmployeeID     uint    `json:"employee_id" form:"employee_id" gorm:"column:employee_id;type:int(10) unsigned;not null;default:0;comment:员工ID"`
	CompletedAt    string  `json:"completed_at" form:"completed_at" gorm:"column:completed_at;type:bigint(20);not null;default:0;comment:完成时间"`
	orm.SoftDeletes
}

func (ClothTailorCuttingPieceProcess) TableName() string {
	return "cloth_tailor_cutting_piece_process"
}
