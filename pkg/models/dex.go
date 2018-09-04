package models

import "encoding/json"

// Dex is a model representing a dex.
type Dex struct {
	tableName struct{} `sql:"dexes,alias:dexes"`

	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
	Shiny    bool   `sql:",notnull" json:"shiny"`
	GameID   string `json:"-"`
	Game     Game   `json:"game"`
	Regional bool   `sql:",notnull" json:"regional"`
}

// MarshalJSON is used to satisfy the Marshaler interface to customize the way
// it's serialized into JSON.
func (d Dex) MarshalJSON() ([]byte, error) {
	total := d.Game.GameFamily.NationalTotal
	if d.Regional {
		total = d.Game.GameFamily.RegionalTotal
	}

	type alias Dex
	return json.Marshal(&struct {
		alias
		Caught int `json:"caught"` // TODO: actually get caught amount
		Total  int `json:"total"`
	}{
		alias: (alias)(d),
		Total: total,
	})
}
