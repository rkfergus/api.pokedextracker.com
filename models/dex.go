package models

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
