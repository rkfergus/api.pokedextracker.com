package games

import "github.com/pokedextracker/api.pokedextracker.com/gamefamilies"

type Game struct {
	ID           string                  `gorm:"primary_key" json:"id"`
	Name         string                  `json:"name"`
	GameFamilyID string                  `json:"-"`
	GameFamily   gamefamilies.GameFamily `gorm:"foreignkey:GameFamilyID" json:"game_family"`
	Order        int                     `json:"order"`
}
