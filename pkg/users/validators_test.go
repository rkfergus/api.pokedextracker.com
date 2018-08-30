package users

import "github.com/bluele/factory-go/factory"

var (
	friendCode = "1234-1234-1234"
	shiny      = false
	regional   = false
)

var createParamsFactory = factory.NewFactory(&createParams{
	Password:   "password",
	FriendCode: &friendCode,
	Title:      "Living Dex",
	Shiny:      &shiny,
	Game:       "red",
	Regional:   &regional,
}).SeqString("Username", func(n string) (interface{}, error) {
	return n, nil
})
