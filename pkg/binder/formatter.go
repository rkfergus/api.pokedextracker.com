package binder

import (
	"fmt"
	"unicode"

	validator "gopkg.in/go-playground/validator.v9"
)

const (
	friendcode = "friendcode"
	max        = "max"
	min        = "min"
	required   = "required"
	token      = "token"
)

func format(err validator.FieldError) string {
	field := toSnake(err.Field())

	switch err.Tag() {
	case friendcode:
		return fmt.Sprintf("%s must be a 12-digit number", field)
	case max:
		return fmt.Sprintf("%s length must be less than or equal to %s characters long", field, err.Param())
	case min:
		return fmt.Sprintf("%s length must be greater than or equal to %s characters long", field, err.Param())
	case required:
		return fmt.Sprintf("%s is required", field)
	case token:
		return fmt.Sprintf("%s must only contain alpha-numeric and underscore characters", field)
	default:
		fmt.Println("actual tag", err.ActualTag())
		fmt.Println("field", field)
		fmt.Println("param", err.Param())
		fmt.Println("struct field", err.StructField())
		fmt.Println("struct namspace", err.StructNamespace())
		fmt.Println("tag", err.Tag())
		fmt.Println("kind", err.Kind())
		fmt.Println("type", err.Type())

		return "NOT IMPLEMENTED YET"
	}
}

func toSnake(in string) string {
	runes := []rune(in)

	var out []rune
	for i := 0; i < len(runes); i++ {
		if i > 0 && (unicode.IsUpper(runes[i]) || unicode.IsNumber(runes[i])) && ((i+1 < len(runes) && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}
