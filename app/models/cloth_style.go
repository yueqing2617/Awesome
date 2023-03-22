package models

import (
	"encoding/json"
	"github.com/goravel/framework/database/orm"
	"strings"
)

type ClothStyle struct {
	orm.Model
	Name         string      `gorm:"column:name;type:varchar(255);not null;default:'';comment:'名称'"`
	Code         string      `gorm:"column:code;type:varchar(255);not null;unique;default:'';comment:'编码'"`
	Picture      string      `gorm:"column:picture;type:text;comment:'图片'"`
	Colors       string      `gorm:"column:colors;type:text;comment:'颜色'"`
	Sizes        string      `gorm:"column:sizes;type:text;comment:'尺寸'"`
	Year         string      `gorm:"column:year;type:varchar(255);default:'';comment:'年份'"`
	Season       string      `gorm:"column:season;type:varchar(255);default:'';comment:'季节'"`
	UnitPrice    float64     `gorm:"column:unit_price;type:decimal(10,2);default:0.00;comment:'单价'"`
	ProcedureStr string      `gorm:"column:procedures;type:json;comment:'工序'"`
	Procedure    []Procedure `gorm:"-"`
	Remark       string      `gorm:"column:remark;type:varchar(255);default:'';comment:'备注'"`
	orm.SoftDeletes
}

func (ClothStyle) TableName() string {
	return "cloth_styles"
}

type Procedure struct {
	Name        string  `json:"name" form:"name"`
	Sort        uint    `json:"sort" form:"sort"`
	Price       float64 `json:"price" form:"price"`
	IsCompleted bool    `json:"is_completed" form:"is_completed"`
}

type ClothStyleResult struct {
	ID         uint        `json:"id"`
	Name       string      `json:"name"`
	Code       string      `json:"code"`
	Picture    string      `json:"picture"`
	Colors     []string    `json:"colors"`
	Sizes      []string    `json:"sizes"`
	Year       string      `json:"year"`
	Season     string      `json:"season"`
	UnitPrice  float64     `json:"unit_price"`
	Procedures []Procedure `json:"procedures"`
	Remark     string      `json:"remark"`
}

func (c *ClothStyle) Value(data *[]ClothStyle) []ClothStyleResult {
	var result []ClothStyleResult
	for _, v := range *data {
		var procedures []Procedure
		_ = json.Unmarshal([]byte(v.ProcedureStr), &procedures)
		result = append(result, ClothStyleResult{
			ID:         v.ID,
			Name:       v.Name,
			Code:       v.Code,
			Picture:    v.Picture,
			Colors:     stringToArray(v.Colors),
			Sizes:      stringToArray(v.Sizes),
			Year:       v.Year,
			Season:     v.Season,
			UnitPrice:  v.UnitPrice,
			Procedures: procedures,
			Remark:     v.Remark,
		})
	}
	return result
}

func stringToArray(str string) []string {
	return strings.Split(str, ",")
}
