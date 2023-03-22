package rules

import (
	"github.com/goravel/framework/contracts/validation"
	"regexp"
)

type Phone struct {
}

// Signature The name of the rule.
func (receiver *Phone) Signature() string {
	return "phone"
}

// Passes Determine if the validation rule passes.
func (receiver *Phone) Passes(data validation.Data, val any, options ...any) bool {
	// 正则验证value是否为手机号
	reg := `^1[3|4|5|6|7|8|9][0-9]{9}$`
	return regexp.MustCompile(reg).MatchString(val.(string))
}

// Message Get the validation error message.
func (receiver *Phone) Message() string {
	return ""
}
