package rules

import (
	"github.com/goravel/framework/contracts/validation"
)

type MaxLength struct {
}

// Signature The name of the rule.
func (receiver *MaxLength) Signature() string {
	return "max_length"
}

// Passes Determine if the validation rule passes.
func (receiver *MaxLength) Passes(data validation.Data, val any, options ...any) bool {
	max := options[0].(int)
	value, is := data.Get(val.(string))
	// value 为空或者不存在也是通过的
	if !is || value == nil {
		return true
	} else {
		return len(value.(string)) <= max
	}
}

// Message Get the validation error message.
func (receiver *MaxLength) Message() string {
	return ""
}
