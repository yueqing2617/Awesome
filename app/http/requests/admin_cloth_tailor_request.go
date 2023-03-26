package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AdminClothTailorRequest struct {
	Name string `form:"name" json:"name"`
}

func (r *AdminClothTailorRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminClothTailorRequest) Rules(ctx http.Context) map[string]string {
	method := ctx.Request().Method()
	switch method {
	case "POST":
		return map[string]string{}
	case "PUT":
		return map[string]string{}
	default:
		return nil
	}
}

func (r *AdminClothTailorRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AdminClothTailorRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AdminClothTailorRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
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
