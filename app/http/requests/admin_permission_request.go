package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AdminPermissionRequest struct {
	Name string `form:"name" json:"name"`
}

func (r *AdminPermissionRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AdminPermissionRequest) Rules() map[string]string {
	return map[string]string{}
}

func (r *AdminPermissionRequest) Messages() map[string]string {
	return map[string]string{}
}

func (r *AdminPermissionRequest) Attributes() map[string]string {
	return map[string]string{}
}

func (r *AdminPermissionRequest) PrepareForValidation(data validation.Data) error {
	return nil
}
