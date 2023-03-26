package models

import (
	"github.com/goravel/framework/database/orm"
)

type ClothOrder struct {
	orm.Model
	Name                string               `gorm:"column:name;type:varchar(255);not null;default:'';comment:'名称'"`
	Code                string               `gorm:"column:code;type:varchar(255);not null;unique;default:'';comment:'订单号'"`
	Customer            *BasicCustomer       `gorm:"foreignKey:CustomerID;references:ID"`
	CustomerID          uint                 `gorm:"column:customer_id;type:bigint(20);not null;default:0;comment:'客户ID'"`
	DeliveryDate        string               `gorm:"column:delivery_date;type:bigint(20);not null;default:0;comment:'交货日期'"`
	OrderType           string               `gorm:"column:order_type;type:varchar(255);not null;default:'';comment:'订单类型'"`
	SalesmanID          uint                 `gorm:"column:salesman_id;type:bigint(20);default:0;comment:'业务员'"`
	Salesman            *BasicSalesman       `gorm:"foreignKey:SalesmanID;references:ID"`
	ClothStyleCode      string               `gorm:"column:cloth_style_code;type:varchar(255);default:'';comment:'款号'"`
	ClothStyleName      string               `gorm:"column:cloth_style_name;type:varchar(255);default:'';comment:'款式名称'"`
	ClothStylePicture   string               `gorm:"column:cloth_style_picture;type:text;default:'';comment:'款式图片'"`
	ClothStyleColors    string               `gorm:"column:cloth_style_colors;type:text;default:'';comment:'款式颜色'"`
	ClothStyleSizes     string               `gorm:"column:cloth_style_sizes;type:text;default:'';comment:'款式尺寸'"`
	ClothStyleYear      string               `gorm:"column:cloth_style_year;type:varchar(255);default:'';comment:'款式年份'"`
	ClothStyleSeason    string               `gorm:"column:cloth_style_season;type:varchar(255);default:'';comment:'款式季节'"`
	ClothStyleUnitPrice float64              `gorm:"column:cloth_style_unit_price;type:decimal(10,2);default:0.00;comment:'款式单价'"`
	Total               uint                 `gorm:"column:total;type:bigint(20);not null;default:0;comment:'总数'"`
	TotalPrice          float64              `gorm:"column:total_price;type:decimal(10,2);not null;default:0.00;comment:'总价'"`
	ContainsStr         string               `gorm:"column:contains;type:text;comment:'包含'"`
	Contains            []ClothOrderContains `gorm:"-"`
	ProceduresStr       string               `gorm:"column:procedures;type:text;comment:'工序'"`
	Procedures          []Procedure          `gorm:"-"`
	Status              int                  `gorm:"column:status;type:tinyint(1);not null;default:0;comment:'状态'"` // 1：未完成 2：已完成 3：已关闭
	Remark              string               `gorm:"column:remark;type:varchar(255);default:'';comment:'备注'"`
	orm.SoftDeletes
}

func (ClothOrder) TableName() string {
	return "cloth_orders"
}

type ClothOrderContains struct {
	Size  string `json:"size" form:"size"`
	Color string `json:"color" form:"color"`
	Num   uint   `json:"num" form:"num"`
}

type ClothOrderResult struct {
	ID                  uint                 `json:"id"`
	Name                string               `json:"name"`
	Code                string               `json:"code"`
	CustomerID          uint                 `json:"customer_id"`
	Customer            *BasicCustomer       `json:"customer"`
	DeliveryDate        string               `json:"delivery_date"`
	OrderType           string               `json:"order_type"`
	SalesmanID          uint                 `json:"salesman_id"`
	Salesman            *BasicSalesman       `json:"salesman"`
	ClothStyleCode      string               `json:"cloth_style_code"`
	ClothStyleName      string               `json:"cloth_style_name"`
	ClothStylePicture   string               `json:"cloth_style_picture"`
	ClothStyleColors    []string             `json:"cloth_style_colors"`
	ClothStyleSizes     []string             `json:"cloth_style_sizes"`
	ClothStyleYear      string               `json:"cloth_style_year"`
	ClothStyleSeason    string               `json:"cloth_style_season"`
	ClothStyleUnitPrice float64              `json:"cloth_style_unit_price"`
	Total               uint                 `json:"total"`
	TotalPrice          float64              `json:"total_price"`
	Contains            []ClothOrderContains `json:"contains"`
	Procedures          []Procedure          `json:"procedures"`
	Remark              string               `json:"remark"`
}
