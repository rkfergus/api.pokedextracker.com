package binder

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	ut "github.com/go-playground/universal-translator"
)

type mockFieldError struct {
	tag   string
	field string
	param string
}

func (e *mockFieldError) Tag() string                       { return e.tag }
func (e *mockFieldError) ActualTag() string                 { return e.tag }
func (e *mockFieldError) Namespace() string                 { return "" }
func (e *mockFieldError) StructNamespace() string           { return "" }
func (e *mockFieldError) Field() string                     { return e.field }
func (e *mockFieldError) StructField() string               { return "" }
func (e *mockFieldError) Value() interface{}                { return "" }
func (e *mockFieldError) Param() string                     { return e.param }
func (e *mockFieldError) Kind() reflect.Kind                { return reflect.TypeOf("").Kind() }
func (e *mockFieldError) Type() reflect.Type                { return reflect.TypeOf("") }
func (e *mockFieldError) Translate(ut ut.Translator) string { return "" }

func TestFormatter(t *testing.T) {
	cases := []struct {
		tag, param, msg string
	}{
		{friendcode, "", "multi_word must be a 12-digit number"},
		{max, "20", "multi_word length must be less than or equal to 20 characters long"},
		{min, "20", "multi_word length must be greater than or equal to 20 characters long"},
		{required, "", "multi_word is required"},
		{token, "", "multi_word must only contain alpha-numeric and underscore characters"},
		{"foo", "", "NOT IMPLEMENTED YET"},
	}

	for _, tt := range cases {
		err := mockFieldError{tt.tag, "MultiWord", tt.param}
		msg := format(&err)
		assert.Equal(t, tt.msg, msg)
	}
}
