package validators

import (
	"reflect"
	"regexp"

	validator "gopkg.in/go-playground/validator.v8"
)

var tokenRE = regexp.MustCompile(`^\w+$`)

func validateToken(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	str, ok := field.Interface().(string)
	if !ok {
		return false
	}
	return tokenRE.MatchString(str)
}
