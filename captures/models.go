package captures

import "github.com/pokedextracker/api.pokedextracker.com/pokemon"

type Capture struct {
	tableName struct{} `sql:"captures,alias:captures"`

	DexID     int             `sql:",pk" json:"dex_id"`
	PokemonID int             `sql:",pk" json:"-"`
	Pokemon   pokemon.Pokemon `json:"pokemon"`
	Captured  bool            `json:"captured"`
}
