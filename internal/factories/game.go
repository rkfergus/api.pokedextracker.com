package factories

import (
	"fmt"

	"github.com/bluele/factory-go/factory"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/models"
)

var (
	defaultGame = &models.Game{
		Name: "Red",
	}
)

// Game is a factory to create fake games for tests.
var Game = factory.NewFactory(defaultGame).
	SeqInt("ID", func(n int) (interface{}, error) {
		return fmt.Sprintf("red_%d", n), nil
	}).
	SeqInt("Order", func(n int) (interface{}, error) {
		return n, nil
	})
