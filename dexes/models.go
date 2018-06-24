package dexes

import (
	"github.com/pokedextracker/api.pokedextracker.com/games"
)

type Dex struct {
	ID       int        `gorm:"primary_key" json:"id"`
	UserID   int        `json:"user_id"`
	Title    string     `json:"title"`
	Slug     string     `json:"slug"`
	Shiny    bool       `json:"shiny"`
	GameID   string     `json:"-"`
	Game     games.Game `gorm:"foreignkey:GameID" json:"game"`
	Regional bool       `json:"regional"`
}
