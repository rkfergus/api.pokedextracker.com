package captures

import "github.com/pokedextracker/api.pokedextracker.com/pokemon"

type Capture struct {
	DexID     int             `gorm:"primary_key" json:"dex_id"`
	PokemonID int             `gorm:"primary_key" json:"-"`
	Pokemon   pokemon.Pokemon `gorm:"foreignkey:PokemonID" json:"pokemon"`
	Captured  bool            `json:"captured"`
}
