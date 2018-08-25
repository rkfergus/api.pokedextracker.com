package models

import (
	"strings"

	"github.com/go-pg/pg/orm"
)

type Pokemon struct {
	tableName struct{} `sql:"pokemon,alias:pokemon"`

	ID           int        `json:"id"`
	NationalID   int        `json:"national_id"`
	Name         string     `json:"name"`
	GameFamilyID string     `json:"-"`
	GameFamily   GameFamily `json:"game_family"`
	Form         *string    `json:"form"`
	Box          *string    `json:"box"`

	RedBlueID                *int `sql:"-" json:"red_blue_id,omitempty"`
	YellowID                 *int `sql:"-" json:"yellow_id,omitempty"`
	GoldSilverID             *int `sql:"-" json:"gold_silver_id,omitempty"`
	CrystalID                *int `sql:"-" json:"crystal_id,omitempty"`
	RubySapphireID           *int `sql:"-" json:"ruby_sapphire_id,omitempty"`
	FireRedLeafGreenID       *int `sql:"-" json:"fire_red_leaf_green_id,omitempty"`
	EmeraldID                *int `sql:"-" json:"emerald_id,omitempty"`
	DiamondPearlID           *int `sql:"-" json:"diamond_pearl_id,omitempty"`
	PlatinumID               *int `sql:"-" json:"platinum_id,omitempty"`
	HeartGoldSoulSilverID    *int `sql:"-" json:"heart_gold_soul_silver_id,omitempty"`
	BlackWhiteID             *int `sql:"-" json:"black_white_id,omitempty"`
	Black2White2ID           *int `sql:"-" json:"black_2_white_2_id,omitempty"`
	XYID                     *int `sql:"-" json:"x_y_id,omitempty"`
	OmegaRubyAlphaSapphireID *int `sql:"-" json:"omega_ruby_alpha_sapphire_id,omitempty"`
	SunMoonID                *int `sql:"-" json:"sun_moon_id,omitempty"`
	UltraSunUltraMoonID      *int `sql:"-" json:"ultra_sun_ultra_moon_id,omitempty"`

	XLocation    string `json:"-"`
	YLocation    string `json:"-"`
	ORLocation   string `json:"-"`
	ASLocation   string `json:"-"`
	SunLocation  string `json:"-"`
	MoonLocation string `json:"-"`
	USLocation   string `json:"-"`
	UMLocation   string `json:"-"`

	XLocations    []string `sql:"-" json:"x_locations"`
	YLocations    []string `sql:"-" json:"y_locations"`
	ORLocations   []string `sql:"-" json:"or_locations"`
	ASLocations   []string `sql:"-" json:"as_locations"`
	SunLocations  []string `sql:"-" json:"sun_locations"`
	MoonLocations []string `sql:"-" json:"moon_locations"`
	USLocations   []string `sql:"-" json:"us_locations"`
	UMLocations   []string `sql:"-" json:"um_locations"`

	EvolutionFamilyID int             `json:"-"`
	EvolutionFamily   EvolutionFamily `sql:"-" json:"evolution_family"`
}

func (p *Pokemon) AfterQuery(db orm.DB) error {
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

func (p *Pokemon) LoadDexNumbers(dexNumbers []*GameFamilyDexNumber) {
	for _, dn := range dexNumbers {
		if dn.PokemonID != p.ID {
			continue
		}
		n := dn.DexNumber
		switch dn.GameFamilyID {
		case RedBlueID:
			p.RedBlueID = &n
		case YellowID:
			p.YellowID = &n
		case GoldSilverID:
			p.GoldSilverID = &n
		case CrystalID:
			p.CrystalID = &n
		case RubySapphireID:
			p.RubySapphireID = &n
		case FireRedLeafGreenID:
			p.FireRedLeafGreenID = &n
		case EmeraldID:
			p.EmeraldID = &n
		case DiamondPearlID:
			p.DiamondPearlID = &n
		case PlatinumID:
			p.PlatinumID = &n
		case HeartGoldSoulSilverID:
			p.HeartGoldSoulSilverID = &n
		case BlackWhiteID:
			p.BlackWhiteID = &n
		case Black2White2ID:
			p.Black2White2ID = &n
		case XYID:
			p.XYID = &n
		case OmegaRubyAlphaSapphireID:
			p.OmegaRubyAlphaSapphireID = &n
		case SunMoonID:
			p.SunMoonID = &n
		case UltraSunUltraMoonID:
			p.UltraSunUltraMoonID = &dn.DexNumber
		}
	}
}

func (p *Pokemon) LoadEvolutions(evolutions []*Evolution) {
	p.EvolutionFamily = EvolutionFamily{[][]PokemonSummary{}, [][]Evolution{}}
	for _, e := range evolutions {
		if e.EvolutionFamilyID != p.EvolutionFamilyID {
			continue
		}

		i := e.Stage - 1
		breed := e.Trigger == "breed"

		var first, second PokemonSummary

		if breed {
			first, second = e.EvolvedPokemon, e.EvolvingPokemon
		} else {
			first, second = e.EvolvingPokemon, e.EvolvedPokemon
		}

		if len(p.EvolutionFamily.Pokemon) <= i {
			p.EvolutionFamily.Pokemon = append(p.EvolutionFamily.Pokemon, []PokemonSummary{})
		}
		if len(p.EvolutionFamily.Pokemon) <= i+1 && len(p.EvolutionFamily.Pokemon) < 3 {
			p.EvolutionFamily.Pokemon = append(p.EvolutionFamily.Pokemon, []PokemonSummary{})
		}

		if !findPokemon(p.EvolutionFamily.Pokemon[i], first) {
			p.EvolutionFamily.Pokemon[i] = append(p.EvolutionFamily.Pokemon[i], first)
		}
		if !findPokemon(p.EvolutionFamily.Pokemon[i+1], second) {
			p.EvolutionFamily.Pokemon[i+1] = append(p.EvolutionFamily.Pokemon[i+1], second)
		}

		if len(p.EvolutionFamily.Evolutions) <= i {
			p.EvolutionFamily.Evolutions = append(p.EvolutionFamily.Evolutions, []Evolution{})
		}
		p.EvolutionFamily.Evolutions[i] = append(p.EvolutionFamily.Evolutions[i], *e)
	}

	// for pokemon without any evolutions, add itself to the Pokemon slice
	if len(p.EvolutionFamily.Pokemon) == 0 {
		summary := PokemonSummary{
			ID:         p.ID,
			NationalID: p.NationalID,
			Name:       p.Name,
			Form:       p.Form,
		}
		p.EvolutionFamily.Pokemon = append(p.EvolutionFamily.Pokemon, []PokemonSummary{summary})
	}
}

type EvolutionFamily struct {
	Pokemon    [][]PokemonSummary `json:"pokemon"`
	Evolutions [][]Evolution      `json:"evolutions"`
}

type PokemonSummary struct {
	tableName struct{} `sql:"pokemon,alias:pokemon"`

	ID         int     `json:"id"`
	NationalID int     `json:"national_id"`
	Name       string  `json:"name"`
	Form       *string `json:"form"`
}

type Evolution struct {
	tableName struct{} `sql:"evolutions,alias:evolutions"`

	EvolvingPokemonID int            `sql:",pk" json:"-"`
	EvolvedPokemonID  int            `sql:",pk" json:"-"`
	EvolvingPokemon   PokemonSummary `json:"-"`
	EvolvedPokemon    PokemonSummary `json:"-"`
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
