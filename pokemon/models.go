package pokemon

import (
	"strings"

	"github.com/pokedextracker/api.pokedextracker.com/dexnumbers"
	"github.com/pokedextracker/api.pokedextracker.com/gamefamilies"
)

type Pokemon struct {
	ID           int                     `gorm:"primary_key" json:"id"`
	NationalID   int                     `json:"national_id"`
	Name         string                  `json:"name"`
	GameFamilyID string                  `json:"-"`
	GameFamily   gamefamilies.GameFamily `gorm:"foreignkey:GameFamilyID" json:"game_family"`
	Form         *string                 `json:"form"`
	Box          *string                 `json:"box"`

	RedBlueID                *int `json:"red_blue_id,omitempty"`
	YellowID                 *int `json:"yellow_id,omitempty"`
	GoldSilverID             *int `json:"gold_silver_id,omitempty"`
	CrystalID                *int `json:"crystal_id,omitempty"`
	RubySapphireID           *int `json:"ruby_sapphire_id,omitempty"`
	FireRedLeafGreenID       *int `json:"fire_red_leaf_green_id,omitempty"`
	EmeraldID                *int `json:"emerald_id,omitempty"`
	DiamondPearlID           *int `json:"diamond_pearl_id,omitempty"`
	PlatinumID               *int `json:"platinum_id,omitempty"`
	HeartGoldSoulSilverID    *int `json:"heart_gold_soul_silver_id,omitempty"`
	BlackWhiteID             *int `json:"black_white_id,omitempty"`
	Black2White2ID           *int `json:"black_2_white_2_id,omitempty"`
	XYID                     *int `json:"x_y_id,omitempty"`
	OmegaRubyAlphaSapphireID *int `json:"omega_ruby_alpha_sapphire_id,omitempty"`
	SunMoonID                *int `json:"sun_moon_id,omitempty"`
	UltraSunUltraMoonID      *int `json:"ultra_sun_ultra_moon_id,omitempty"`

	XLocation    string `json:"-"`
	YLocation    string `json:"-"`
	ORLocation   string `json:"-"`
	ASLocation   string `json:"-"`
	SunLocation  string `json:"-"`
	MoonLocation string `json:"-"`
	USLocation   string `json:"-"`
	UMLocation   string `json:"-"`

	XLocations    []string `gorm:"-" json:"x_locations"`
	YLocations    []string `gorm:"-" json:"y_locations"`
	ORLocations   []string `gorm:"-" json:"or_locations"`
	ASLocations   []string `gorm:"-" json:"as_locations"`
	SunLocations  []string `gorm:"-" json:"sun_locations"`
	MoonLocations []string `gorm:"-" json:"moon_locations"`
	USLocations   []string `gorm:"-" json:"us_locations"`
	UMLocations   []string `gorm:"-" json:"um_locations"`

	EvolutionFamilyID int             `json:"-"`
	EvolutionFamily   EvolutionFamily `gorm:"-" json:"evolution_family"`
}

func (*Pokemon) TableName() string {
	return "pokemon"
}

func (p *Pokemon) AfterFind() error {
	p.XLocations = splitLocations(p.XLocation)
	p.YLocations = splitLocations(p.YLocation)
	p.ORLocations = splitLocations(p.ORLocation)
	p.ASLocations = splitLocations(p.ASLocation)
	p.SunLocations = splitLocations(p.SunLocation)
	p.MoonLocations = splitLocations(p.MoonLocation)
	p.USLocations = splitLocations(p.USLocation)
	p.UMLocations = splitLocations(p.UMLocation)

	return nil
}

func (p *Pokemon) LoadDexNumbers(dexNumbers []*dexnumbers.GameFamilyDexNumber) {
	for _, dn := range dexNumbers {
		if dn.PokemonID != p.ID {
			continue
		}
		n := dn.DexNumber
		switch dn.GameFamilyID {
		case gamefamilies.RedBlueID:
			p.RedBlueID = &n
		case gamefamilies.YellowID:
			p.YellowID = &n
		case gamefamilies.GoldSilverID:
			p.GoldSilverID = &n
		case gamefamilies.CrystalID:
			p.CrystalID = &n
		case gamefamilies.RubySapphireID:
			p.RubySapphireID = &n
		case gamefamilies.FireRedLeafGreenID:
			p.FireRedLeafGreenID = &n
		case gamefamilies.EmeraldID:
			p.EmeraldID = &n
		case gamefamilies.DiamondPearlID:
			p.DiamondPearlID = &n
		case gamefamilies.PlatinumID:
			p.PlatinumID = &n
		case gamefamilies.HeartGoldSoulSilverID:
			p.HeartGoldSoulSilverID = &n
		case gamefamilies.BlackWhiteID:
			p.BlackWhiteID = &n
		case gamefamilies.Black2White2ID:
			p.Black2White2ID = &n
		case gamefamilies.XYID:
			p.XYID = &n
		case gamefamilies.OmegaRubyAlphaSapphireID:
			p.OmegaRubyAlphaSapphireID = &n
		case gamefamilies.SunMoonID:
			p.SunMoonID = &n
		case gamefamilies.UltraSunUltraMoonID:
			p.UltraSunUltraMoonID = &dn.DexNumber
		}
	}
}

func (p *Pokemon) LoadEvolutions(evolutions []*Evolution) {
	p.EvolutionFamily = EvolutionFamily{make([][]PokemonSummary, 3), make([][]Evolution, 2)}
	for _, e := range evolutions {
		if e.EvolutionFamilyID != p.EvolutionFamilyID {
			continue
		}

		i := e.Stage - 1
		breed := e.Trigger == "breed"

		var (
			first  PokemonSummary
			second PokemonSummary
		)

		if breed {
			first, second = e.EvolvedPokemon, e.EvolvingPokemon
		} else {
			first, second = e.EvolvingPokemon, e.EvolvedPokemon
		}

		if !findPokemon(p.EvolutionFamily.Pokemon[i], first) {
			p.EvolutionFamily.Pokemon[i] = append(p.EvolutionFamily.Pokemon[i], first)
		}
		if !findPokemon(p.EvolutionFamily.Pokemon[i+1], second) {
			p.EvolutionFamily.Pokemon[i+1] = append(p.EvolutionFamily.Pokemon[i+1], second)
		}

		p.EvolutionFamily.Evolutions[i] = append(p.EvolutionFamily.Evolutions[i], *e)
	}
}

type EvolutionFamily struct {
	Pokemon    [][]PokemonSummary `json:"pokemon"`
	Evolutions [][]Evolution      `json:"evolutions"`
}

type PokemonSummary struct {
	ID         int     `gorm:"primary_key" json:"id"`
	NationalID int     `json:"national_id"`
	Name       string  `json:"name"`
	Form       *string `json:"form"`
}

func (*PokemonSummary) TableName() string {
	return "pokemon"
}

type Evolution struct {
	EvolvingPokemonID int            `gorm:"primary_key" json:"-"`
	EvolvedPokemonID  int            `gorm:"primary_key" json:"-"`
	EvolvingPokemon   PokemonSummary `gorm:"foreignkey:EvolvingPokemonID" json:"-"`
	EvolvedPokemon    PokemonSummary `gorm:"foreignkey:EvolvedPokemonID" json:"-"`
	EvolutionFamilyID int            `json:"-"`
	Stage             int            `json:"-"`
	Trigger           string         `json:"trigger"`
	Level             *int           `json:"level,omitempty"`
	Stone             *string        `json:"stone,omitempty"`
	HeldItem          *string        `json:"held_item,omitempty"`
	Notes             *string        `json:"notes,omitempty"`
}

func splitLocations(location string) []string {
	if location != "" {
		return strings.Split(location, ", ")
	}
	return []string{}
}

func findPokemon(summaries []PokemonSummary, summary PokemonSummary) bool {
	for _, s := range summaries {
		if s.ID == summary.ID {
			return true
		}
	}
	return false
}
