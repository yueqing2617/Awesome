package requests

import (
	"Awesome/app/models"
	"errors"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AdminClothOrderRequest struct {
	Name                string                      `form:"name" json:"name"`
	Code                string                      `form:"code" json:"code"`
	CustomerID          uint                        `form:"customer_id" json:"customer_id"`
	DeliveryDate        string                      `form:"delivery_date" json:"delivery_date"`
	OrderType           string                      `form:"order_type" json:"order_type"`
	SalesmanID          uint                        `form:"salesman_id" json:"salesman_id"`
	ClothStyleCode      string                      `form:"cloth_style_code" json:"cloth_style_code"`
	ClothStyleName      string                      `form:"cloth_style_name" json:"cloth_style_name"`
	ClothStylePicture   string                      `form:"cloth_style_picture" json:"cloth_style_picture"`
	ClothStyleColors    []string                    `form:"cloth_style_colors" json:"cloth_style_colors"`
	ClothStyleSizes     []string                    `form:"cloth_style_sizes" json:"cloth_style_sizes"`
	ClothStyleYear      string                      `form:"cloth_style_year" json:"cloth_style_year"`
	ClothStyleSeason    string                      `form:"cloth_style_season" json:"cloth_style_season"`
	ClothStyleUnitPrice float64                     `form:"cloth_style_unit_price" json:"cloth_style_unit_price"`
	Total               uint                        `form:"total" json:"total"`
	TotalPrice          float64                     `form:"total_price" json:"total_price"`
	Contains            []models.ClothOrderContains `form:"contains" json:"contains"`
	Procedures          []models.Procedure          `form:"procedures" json:"procedures"`
	Remark              string                      `form:"remark" json:"remark"`
}

func (r *AdminClothOrderRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminClothOrderRequest) Rules(ctx http.Context) map[string]string {

	return map[string]string{
		"name":                   "required|min_len:2|max_len:255",
		"code":                   "required|min_len:2|max_len:255",
		"customer_id":            "required|numeric",
		"delivery_date":          "required",
		"order_type":             "required",
		"salesman_id":            "required|numeric",
		"cloth_style_code":       "required",
		"cloth_style_name":       "required|min_len:2|max_len:255",
		"cloth_style_picture":    "required|min_len:1",
		"cloth_style_colors":     "required|min_len:1",
		"cloth_style_sizes":      "required|min_len:1",
		"cloth_style_year":       "required",
		"cloth_style_season":     "required",
		"cloth_style_unit_price": "required|numeric",
		"total":                  "required|numeric",
		"total_price":            "required|numeric",
		"contains":               "required|slice",
		"procedures":             "required|slice",
		"remark":                 "required",
	}
}

func (r *AdminClothOrderRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"name.required":                   "名称不能为空",
		"name.min_len":                    "名称不能少于2个字符",
		"name.max_len":                    "名称不能多于255个字符",
		"code.required":                   "编号不能为空",
		"code.min_len":                    "编号不能少于2个字符",
		"code.max_len":                    "编号不能多于255个字符",
		"customer_id.required":            "客户不能为空",
		"customer_id.numeric":             "客户必须是数字",
		"delivery_date.required":          "交货日期不能为空",
		"order_type.required":             "订单类型不能为空",
		"salesman_id.required":            "业务员不能为空",
		"salesman_id.numeric":             "业务员必须是数字",
		"cloth_style_code.required":       "款式编号不能为空",
		"cloth_style_name.required":       "款式名称不能为空",
		"cloth_style_name.min_len":        "款式名称不能少于2个字符",
		"cloth_style_name.max_len":        "款式名称不能多于255个字符",
		"cloth_style_picture.required":    "款式图片不能为空",
		"cloth_style_picture.min_len":     "款式图片不能少于1个字符",
		"cloth_style_colors.required":     "款式颜色不能为空",
		"cloth_style_colors.min_len":      "款式颜色不能少于1个字符",
		"cloth_style_sizes.required":      "款式尺寸不能为空",
		"cloth_style_sizes.min_len":       "款式尺寸不能少于1个字符",
		"cloth_style_year.required":       "款式年份不能为空",
		"cloth_style_season.required":     "款式季节不能为空",
		"cloth_style_unit_price.required": "款式单价不能为空",
		"cloth_style_unit_price.numeric":  "款式单价必须是数字",
		"total.required":                  "总数不能为空",
		"total.numeric":                   "总数必须是数字",
		"total_price.required":            "总价不能为空",
		"total_price.numeric":             "总价必须是数字",
		"contains.required":               "包含不能为空",
		"contains.slice":                  "包含必须是切片",
		"procedures.required":             "工序不能为空",
		"procedures.slice":                "工序必须是切片",
		"remark.required":                 "备注不能为空",
	}
}

func (r *AdminClothOrderRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"name":                   "名称",
		"code":                   "编号",
		"customer_id":            "客户",
		"delivery_date":          "交货日期",
		"order_type":             "订单类型",
		"salesman":               "业务员",
		"cloth_style_code":       "款式编号",
		"cloth_style_name":       "款式名称",
		"cloth_style_picture":    "款式图片",
		"cloth_style_colors":     "款式颜色",
		"cloth_style_sizes":      "款式尺寸",
		"cloth_style_year":       "款式年份",
		"cloth_style_season":     "款式季节",
		"cloth_style_unit_price": "款式单价",
		"total":                  "总数",
		"total_price":            "总价",
		"contains":               "包含",
		"procedures":             "工序",
		"remark":                 "备注",
	}
}

func (r *AdminClothOrderRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	method := ctx.Request().Method()
	switch method {
	case "POST":
		return nil
	case "PUT":
		id := ctx.Request().Input("id")
		if id == "" {
			return errors.New("id不能为空")
		}
		return nil
	default:
		return nil
	}
}
