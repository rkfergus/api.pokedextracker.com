package gamefamilies

type GameFamily struct {
	tableName struct{} `sql:"game_families,alias:game_families"`

	ID            string `json:"id"`
	Generation    int    `json:"generation"`
	RegionalTotal int    `json:"regional_total"`
	NationalTotal int    `json:"national_total"`
	Order         int    `json:"order"`
	Published     bool   `json:"published"`
}

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
