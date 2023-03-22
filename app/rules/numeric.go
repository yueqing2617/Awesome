package rules

import (
	"github.com/goravel/framework/contracts/validation"
	"regexp"
)

type Numeric struct {
}

// Signature The name of the rule.
func (receiver *Numeric) Signature() string {
	return "numeric"
}

// Passes Determine if the validation rule passes.
func (receiver *Numeric) Passes(data validation.Data, val any, options ...any) bool {
	// 正则表达式判断val 是否为数字
	reg := regexp.MustCompile(`^[0-9]+$`)
	return reg.MatchString(val.(string))
}

// Message Get the validation error message.
func (receiver *Numeric) Message() string {
	return ""
}
