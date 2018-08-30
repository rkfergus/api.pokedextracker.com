package binder

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	validator "gopkg.in/go-playground/validator.v9"
)

type mockFieldLevel struct {
	value string
}

func (fl *mockFieldLevel) Top() reflect.Value      { return reflect.ValueOf("") }
func (fl *mockFieldLevel) Parent() reflect.Value   { return reflect.ValueOf("") }
func (fl *mockFieldLevel) Field() reflect.Value    { return reflect.ValueOf(fl.value) }
func (fl *mockFieldLevel) FieldName() string       { return "" }
func (fl *mockFieldLevel) StructFieldName() string { return "" }
func (fl *mockFieldLevel) Param() string           { return "" }
func (fl *mockFieldLevel) ExtractType(field reflect.Value) (reflect.Value, reflect.Kind, bool) {
	return reflect.ValueOf(""), reflect.TypeOf("").Kind(), false
}
func (fl *mockFieldLevel) GetStructFieldOK() (reflect.Value, reflect.Kind, bool) {
	return reflect.ValueOf(""), reflect.TypeOf("").Kind(), false
}

func TestValidatorFunctions(t *testing.T) {
	cases := []struct {
		value  string
		fn     validator.Func
		passed bool
		msg    string
	}{
		{"", friendcodeFunc, true, "friendcode: empty strings should be allowed"},
		{"1234-1234-1234", friendcodeFunc, true, "friendcode: 12-digit codes should be allowed"},
		{"foo", friendcodeFunc, false, "friendcode: others should not be allowed"},
		{"aA0_", tokenFunc, true, "token: alphanum and underscores should be allowed"},
		{"-", tokenFunc, false, "token: others should not be allowed"},
	}

	for _, tt := range cases {
		fl := mockFieldLevel{tt.value}
		passed := tt.fn(&fl)
		assert.Equal(t, tt.passed, passed, tt.msg)
	}
}
