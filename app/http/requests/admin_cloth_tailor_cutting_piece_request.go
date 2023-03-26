package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type Piece struct {
	Layer uint   `form:"layer" json:"layer"`
	Color string `form:"color" json:"color"`
	Size  string `form:"size" json:"size"`
	Num   int    `form:"num" json:"num"`
}

type AdminClothTailorCuttingPieceRequest struct {
	BedNum   string  `form:"bed_num" json:"bed_num"`
	StartNum uint    `form:"start_num" json:"start_num"`
	Pieces   []Piece `form:"pieces" json:"pieces"`
}

func (r *AdminClothTailorCuttingPieceRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminClothTailorCuttingPieceRequest) Rules(ctx http.Context) map[string]string {
	method := ctx.Request().Method()
	switch method {
	case "POST":
		return map[string]string{
			"bed_num":   "required|string:1,10",
			"start_num": "required|numeric|min:1",
			"pieces":    "required|slice|min_len:1",
		}
	case "PUT":
		return map[string]string{
			"bed_num":   "required|string:1,10",
			"start_num": "required|numeric|min:1",
			"pieces":    "required|slice|min_len:1",
		}
	default:
		return nil
	}
}

func (r *AdminClothTailorCuttingPieceRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"bed_num.required":   "床号不能为空",
		"bed_num.string":     "床号必须为字符串",
		"bed_num.min":        "床号长度不能小于1",
		"bed_num.max":        "床号长度不能大于10",
		"start_num.required": "起始扎号不能为空",
		"start_num.numeric":  "起始扎号必须为数字",
		"start_num.min":      "起始扎号不能小于1",
		"pieces.required":    "裁剪件数不能为空",
		"pieces.slice":       "裁剪件数必须为切片",
	}
}

func (r *AdminClothTailorCuttingPieceRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"bed_num":   "床号",
		"start_num": "起始扎号",
		"pieces":    "裁剪件数",
	}
}

func (r *AdminClothTailorCuttingPieceRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	method := ctx.Request().Method()
	switch method {
	case "POST":
		return nil
	case "PUT":
		return nil
	default:
		return nil
	}
}
