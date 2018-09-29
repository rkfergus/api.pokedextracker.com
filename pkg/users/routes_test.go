package users

import (
	"testing"

	"github.com/labstack/echo"
	"github.com/pokedextracker/api.pokedextracker.com/internal/test"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/application"
	"github.com/stretchr/testify/assert"
)

func TestRegisterRoutes(t *testing.T) {
	e := echo.New()
	app := &application.App{}

	RegisterRoutes(e, app)

	routes := test.FilterRoutes(e.Routes())
	assert.Len(t, routes, 4)
}
