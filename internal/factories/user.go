package factories

import (
	"github.com/bluele/factory-go/factory"
	"github.com/pokedextracker/api.pokedextracker.com/pkg/models"
)

var (
	friendCode = "1234-1234-1234"
)

// User is a factory to create fake users for tests.
var User = factory.NewFactory(&models.User{
	Password:   "password",
	FriendCode: &friendCode,
}).SeqString("Username", func(n string) (interface{}, error) {
	return n, nil
})
