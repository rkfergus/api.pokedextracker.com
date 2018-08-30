package binder

import (
	"regexp"

	validator "gopkg.in/go-playground/validator.v9"
)

var (
	friendcodeRE = regexp.MustCompile(`^\d{4}-\d{4}-\d{4}$`)
	tokenRE      = regexp.MustCompile(`^\w+$`)
)

func friendcodeFunc(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	return str == "" || friendcodeRE.MatchString(str)
}

func tokenFunc(fl validator.FieldLevel) bool {
	return tokenRE.MatchString(fl.Field().String())
}
