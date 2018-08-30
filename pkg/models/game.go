package models

type Game struct {
	tableName struct{} `sql:"games,alias:games"`

	ID           string     `json:"id"`
	Name         string     `json:"name"`
	GameFamilyID string     `json:"-"`
	GameFamily   GameFamily `json:"game_family"`
	Order        int        `json:"order"`
}
