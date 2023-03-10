package rules

import (
	"github.com/goravel/framework/contracts/validation"
	"regexp"
)

type RegexRule struct {
}

// Signature The name of the rule.
func (receiver *RegexRule) Signature() string {
	return "regex_rule"
}

// Passes Determine if the validation rule passes.
func (receiver *RegexRule) Passes(data validation.Data, val any, options ...any) bool {
	reg := options[0].(string)
	if val != nil {
		return regexp.MustCompile(reg).MatchString(val.(string))
	}
	return false
}

// Message Get the validation error message.
func (receiver *RegexRule) Message() string {
	return ""
}
