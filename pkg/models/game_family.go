package models

// GameFamily is a model representing a game family. A game family is a set of
// games that were released at the same time, usually as complementing games.
type GameFamily struct {
	tableName struct{} `sql:"game_families,alias:game_families"`

	ID            string `json:"id"`
	Generation    int    `json:"generation"`
	RegionalTotal int    `json:"regional_total"`
	NationalTotal int    `json:"national_total"`
	Order         int    `json:"order"`
	Published     bool   `json:"published"`
}

// These constants correspond to the IDs of the game families.
const (
	RedBlueID                = "red_blue"
	YellowID                 = "yellow"
	GoldSilverID             = "gold_silver"
	CrystalID                = "crystal"
	RubySapphireID           = "ruby_sapphire"
	FireRedLeafGreenID       = "fire_red_leaf_green"
	EmeraldID                = "emerald"
	DiamondPearlID           = "diamond_pearl"
	PlatinumID               = "platinum"
	HeartGoldSoulSilverID    = "heart_gold_soul_silver"
	BlackWhiteID             = "black_white"
	Black2White2ID           = "black_2_white_2"
	XYID                     = "x_y"
	OmegaRubyAlphaSapphireID = "omega_ruby_alpha_sapphire"
	SunMoonID                = "sun_moon"
	UltraSunUltraMoonID      = "ultra_sun_ultra_moon"
)
