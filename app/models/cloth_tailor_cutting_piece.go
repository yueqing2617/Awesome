package models

import (
	"github.com/goravel/framework/database/orm"
)

type ClothTailorCuttingPieces struct {
	orm.Model
	ClothTailorID  uint                             `gorm:"column:cloth_tailor_id;type:int(10) unsigned;not null;comment:裁剪订单ID" json:"cloth_tailor_id"`
	ClothTailor    *ClothTailor                     `gorm:"foreignKey:cloth_tailor_id;references:id;" json:"cloth_tailor"`
	ClothStyleCode string                           `gorm:"column:cloth_style_code;type:varchar(255);not null;comment:布料款式编码" json:"cloth_style_code"`
	ClothStyle     *ClothStyle                      `gorm:"foreignKey:cloth_style_code;references:code;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"cloth_style"`
	BedNumber      string                           `gorm:"column:bed;type:varchar(255) unsigned;not null;comment:床次" json:"bed"`
	Number         uint                             `gorm:"column:number;type:int(10) unsigned;not null;comment:扎号" json:"number"`
	Layer          uint                             `gorm:"column:layer;type:int(10) unsigned;not null;comment:拉布层数" json:"layer"`
	Color          string                           `gorm:"column:color;type:varchar(255);not null;comment:颜色" json:"color"`
	Size           string                           `gorm:"column:size;type:varchar(255);not null;comment:尺寸" json:"size"`
	Process        []ClothTailorCuttingPieceProcess `gorm:"foreignKey:cutting_piece_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"cutting_piece_process"`
	IsCompleted    bool                             `gorm:"column:is_completed;type:tinyint(1);not null;comment:是否完成" json:"is_completed"`
}

func (ClothTailorCuttingPieces) TableName() string {
	return "cloth_tailor_cutting_pieces"
}
