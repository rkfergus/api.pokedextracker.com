package dexes

import (
	"github.com/pokedextracker/api.pokedextracker.com/games"
)

type Dex struct {
	tableName struct{} `sql:"dexes,alias:dexes"`

	ID       int        `json:"id"`
	UserID   int        `json:"user_id"`
	Title    string     `json:"title"`
	Slug     string     `json:"slug"`
	Shiny    bool       `json:"shiny"`
	GameID   string     `json:"-"`
	Game     games.Game `json:"game"`
	Regional bool       `json:"regional"`
}
