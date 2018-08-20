package binder

import (
	"regexp"

	validator "gopkg.in/go-playground/validator.v9"
)

var (
	friendCodeRE = regexp.MustCompile(`^\d{4}-\d{4}-\d{4}$`)
	tokenRE      = regexp.MustCompile(`^\w+$`)
)

func friendCode(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	return str == "" || friendCodeRE.MatchString(str)
}

func token(fl validator.FieldLevel) bool {
	return tokenRE.MatchString(fl.Field().String())
}
