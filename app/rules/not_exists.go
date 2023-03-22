package rules

import (
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/facades"
)

type NotExists struct {
}

// Signature The name of the rule.
func (receiver *NotExists) Signature() string {
	return "not_exists"
}

// Passes Determine if the validation rule passes.
func (receiver *NotExists) Passes(data validation.Data, val any, options ...any) bool {
	table := options[0].(string)
	field := options[1].(string)
	var count int64
	facades.Orm.Query().Raw("select * from "+table+" where "+field+" = ?", val).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

// Message Get the validation error message.
func (receiver *NotExists) Message() string {
	return "The :attribute must be equal to :other."
}
