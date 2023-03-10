package rules

import (
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/facades"
)

type Unique struct {
}

// Signature The name of the rule.
func (receiver *Unique) Signature() string {
	return "unique"
}

// Passes Determine if the validation rule passes.
func (receiver *Unique) Passes(data validation.Data, val any, options ...any) bool {
	table := options[0].(string)
	field := options[1].(string)
	if val == nil {
		return true
	}
	var count int64
	err := facades.Orm.Query().Raw("SELECT COUNT(*) FROM "+table+" WHERE "+field+" = ?", val).Count(&count)
	if err != nil {
		return true
	}
	return count == 0
}

// Message Get the validation error message.
func (receiver *Unique) Message() string {
	return "This :attribute already exists in the database."
}
