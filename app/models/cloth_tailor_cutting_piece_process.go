package models

import "github.com/goravel/framework/database/orm"

type ClothTailorCuttingPieceProcess struct {
	orm.Model
	Name           string  `json:"name" form:"name"`
	Sort           uint    `json:"sort" form:"sort"`
	Price          float64 `json:"price" form:"price"`
	CuttingPieceID uint    `gorm:"column:cutting_piece_id;type:int(10) unsigned;not null;comment:裁剪订单ID" json:"cutting_piece_id"`
	IsCompleted    bool    `json:"is_completed" form:"is_completed"`
	EmployeeID     uint    `json:"employee_id" form:"employee_id"`
	CompletedAt    string  `json:"completed_at" form:"completed_at"`
	orm.SoftDeletes
}

func (ClothTailorCuttingPieceProcess) TableName() string {
	return "cloth_tailor_cutting_piece_process"
}
