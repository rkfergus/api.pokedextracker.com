package factories

import (
	"fmt"

	"github.com/bluele/factory-go/factory"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/models"
)

var (
	defaultGameFamily = &models.GameFamily{
		Generation:    1,
		RegionalTotal: 151,
		NationalTotal: 151,
		Published:     false,
	}
)

// GameFamily is a factory to create fake game families for tests.
var GameFamily = factory.NewFactory(defaultGameFamily).
	SeqInt("ID", func(n int) (interface{}, error) {
		return fmt.Sprintf("red_blue_%d", n), nil
	}).
	SeqInt("Order", func(n int) (interface{}, error) {
		return n, nil
	})
