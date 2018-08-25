package users

import (
	"fmt"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/pokedextracker/api.pokedextracker.com/application"
	"github.com/stretchr/testify/assert"
)

func TestRegisterRoutes(t *testing.T) {
	e := echo.New()
	app := &application.App{}

	RegisterRoutes(e, app)

	routes := []string{}

	for _, r := range e.Routes() {
		// only count routes added by the function
		if strings.Contains(r.Name, "Handler") {
			routes = append(routes, fmt.Sprintf("%s %s", r.Method, r.Path))
		}
	}

	assert.Len(t, routes, 4)
}
