package rules

import (
	"Awesome/app/http/helper"
	"github.com/goravel/framework/contracts/validation"
)

type Captcha struct {
}

// Signature The name of the rule.
func (receiver *Captcha) Signature() string {
	return "captcha"
}

// Passes Determine if the validation rule passes.
func (receiver *Captcha) Passes(data validation.Data, val any, options ...any) bool {
	id, _ := data.Get("app_secret")
	if id == nil || id.(string) == "" || val.(string) == "" {
		return false
	}
	return helper.VerifyCaptcha(id.(string), val.(string))
}

// Message Get the validation error message.
func (receiver *Captcha) Message() string {
	return "验证码错误"
}
