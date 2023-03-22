package requests

import (
	"Awesome/app/http/helper"
	"Awesome/app/models"
	"errors"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"strconv"
)

type AdminClothStyleRequest struct {
	Name       string             `form:"name" json:"name"`
	Code       string             `form:"code" json:"code"`
	Picture    string             `form:"picture" json:"picture"`
	Colors     []string           `form:"colors" json:"colors"`
	Sizes      []string           `form:"sizes" json:"sizes"`
	Year       string             `form:"year" json:"year"`
	Season     string             `form:"season" json:"season"`
	UnitPrice  float64            `form:"unit_price" json:"unit_price"`
	Procedures []models.Procedure `form:"procedures" json:"procedures"`
	Remark     string             `form:"remark" json:"remark"`
}

func (r *AdminClothStyleRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminClothStyleRequest) Rules(ctx http.Context) map[string]string {
	method := ctx.Request().Method()
	switch method {
	case "POST":
		return map[string]string{
			"name":       "required|min_len:2|max_len:255",
			"code":       "required|min_len:2|max_len:255|unique:cloth_styles,code",
			"picture":    "max_len:255",
			"colors":     "required|min_len:1",
			"sizes":      "required|min_len:1",
			"year":       "required|min_len:2|max_len:4",
			"season":     "required|min_len:2|max_len:255",
			"unit_price": "required|number",
			"procedures": "required|min_len:1|slice",
			"remark":     "max_len:255",
		}
	case "PUT":
		return map[string]string{
			"name":       "required|min_len:2|max_len:255",
			"picture":    "max_len:255",
			"colors":     "required|min_len:1",
			"sizes":      "required|min_len:1",
			"year":       "required|min_len:2|max_len:4",
			"season":     "required|min_len:2|max_len:255",
			"unit_price": "required|number",
			"procedures": "required|min_len:1|slice",
			"remark":     "max_len:255",
		}
	default:
		return nil
	}
}

func (r *AdminClothStyleRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"name.required":       "名称不能为空",
		"name.min_len":        "名称长度不能小于2",
		"name.max_len":        "名称长度不能大于255",
		"code.required":       "款号不能为空",
		"code.min_len":        "款号长度不能小于2",
		"code.max_len":        "款号长度不能大于255",
		"code.unique":         "款号已存在",
		"code.not_exists":     "款号不能更改",
		"picture.max_len":     "图片长度不能大于255",
		"colors.required":     "颜色不能为空",
		"colors.min_len":      "颜色长度不能小于1",
		"sizes.required":      "尺码不能为空",
		"sizes.min_len":       "尺码长度不能小于1",
		"year.required":       "年份不能为空",
		"year.min_len":        "年份长度不能小于2",
		"year.max_len":        "年份长度不能大于4",
		"season.required":     "季节不能为空",
		"season.min_len":      "季节长度不能小于2",
		"season.max_len":      "季节长度不能大于255",
		"unit_price.required": "单价不能为空",
		"unit_price.number":   "单价必须是数字",
		"procedures.required": "工序不能为空",
		"procedures.min_len":  "工序长度不能小于1",
		"procedures.slice":    "工序必须是切片",
		"remark.max_len":      "备注长度不能大于255",
	}
}

func (r *AdminClothStyleRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"name":       "名称",
		"code":       "款号",
		"picture":    "图片",
		"colors":     "颜色",
		"sizes":      "尺码",
		"year":       "年份",
		"season":     "季节",
		"unit_price": "单价",
		"procedures": "工序",
		"remark":     "备注",
	}
}

func (r *AdminClothStyleRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	method := ctx.Request().Method()
	procedures, _ := data.Get("procedures")
	if procedures != nil {
		procedures := procedures.([]interface{})
		for i := 0; i < len(procedures); i++ {
			name := procedures[i].(map[string]interface{})["name"]
			price := procedures[i].(map[string]interface{})["price"]
			sort := procedures[i].(map[string]interface{})["sort"]
			isCompleted := procedures[i].(map[string]interface{})["is_completed"]
			if name == "" {
				return errors.New("工序[" + strconv.Itoa(i) + "]名称不能为空")
			}
			if price == "" {
				return errors.New("工序[" + strconv.Itoa(i) + "]价格不能为空")
			}
			if sort == "" {
				procedures[i].(map[string]interface{})["sort"] = 0
			}
			if isCompleted == "" {
				procedures[i].(map[string]interface{})["is_completed"] = false
			}
		}
		_ = data.Set("procedures", procedures)
	}
	switch method {
	case "POST":
		code, _ := data.Get("code")
		if code == "" {
			_ = data.Set("code", helper.RandomString(12))
		}
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
