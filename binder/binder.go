package binder

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	mold "gopkg.in/go-playground/mold.v2"
	"gopkg.in/go-playground/mold.v2/modifiers"
	validator "gopkg.in/go-playground/validator.v9"
)

type Binder struct {
	db       *echo.DefaultBinder
	conform  *mold.Transformer
	validate *validator.Validate
}

func New() *Binder {
	db := &echo.DefaultBinder{}
	conform := modifiers.New()
	validate := validator.New()

	validate.RegisterValidation("friendcode", friendCode)
	validate.RegisterValidation("token", token)

	return &Binder{db, conform, validate}
}

func (b *Binder) Bind(i interface{}, c echo.Context) error {
	if err := b.db.Bind(i, c); err != nil {
		fmt.Println(err)
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
