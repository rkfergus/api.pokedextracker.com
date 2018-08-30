package binder

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	mold "gopkg.in/go-playground/mold.v2"
	"gopkg.in/go-playground/mold.v2/modifiers"
	validator "gopkg.in/go-playground/validator.v9"
)

// Binder is a custom struct that implements the Echo Binder interface. It binds
// to a struct, uses mold to clean up the params, and validator to validate
// them.
type Binder struct {
	db       *echo.DefaultBinder
	conform  *mold.Transformer
	validate *validator.Validate
}

// New initializes a new Binder instance with the appropriate validation
// functions registered.
func New() (*Binder, error) {
	db := &echo.DefaultBinder{}
	conform := modifiers.New()
	validate := validator.New()

	if err := validate.RegisterValidation("friendcode", friendcodeFunc); err != nil {
		return nil, err
	}
	if err := validate.RegisterValidation("token", tokenFunc); err != nil {
		return nil, err
	}

	return &Binder{db, conform, validate}, nil
}

// Bind binds, modifies, and validates payloads against the given struct.
func (b *Binder) Bind(i interface{}, c echo.Context) error {
	if err := b.db.Bind(i, c); err != nil {
		return err
	}
	if err := b.conform.Struct(context.Background(), i); err != nil {
		return err
	}
	if err := b.validate.Struct(i); err != nil {
		errs := err.(validator.ValidationErrors)
		msg := format(errs[0])
		return echo.NewHTTPError(http.StatusUnprocessableEntity, msg)
	}
	return nil
}
