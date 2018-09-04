package models

// Capture is a model representing a dex capture. A capture row only exists if
// the Pokemon is captured i.e. Captured should never be false.
type Capture struct {
	tableName struct{} `sql:"captures,alias:captures"`

	DexID     int     `sql:",pk" json:"dex_id"`
	PokemonID int     `sql:",pk" json:"-"`
	Pokemon   Pokemon `json:"pokemon"`
	Captured  bool    `json:"captured"`
}
