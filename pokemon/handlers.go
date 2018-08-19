package pokemon

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pokedextracker/api.pokedextracker.com/application"
	"github.com/pokedextracker/api.pokedextracker.com/dexnumbers"
)

func listHandler(c echo.Context) error {
	app := c.Get("app").(*application.App)

	var pokemon []*Pokemon
	var dexNumbers []*dexnumbers.GameFamilyDexNumber
	var evolutions []*Evolution

	err := app.DB.Model(&pokemon).Relation("GameFamily").Order("national_order ASC").Select()
	if err != nil {
		return err
	}
	err = app.DB.Model(&dexNumbers).Select()
	if err != nil {
		return err
	}
	err = app.DB.
		Model(&evolutions).
		Relation("EvolvingPokemon").
		Relation("EvolvedPokemon").
		Join("INNER JOIN pokemon AS evolved ON evolutions.evolved_pokemon_id = evolved.id").
		Join("INNER JOIN pokemon AS evolving ON evolutions.evolving_pokemon_id = evolving.id").
		OrderExpr("CASE WHEN trigger = 'breed' THEN evolving.national_id ELSE evolved.national_id END, trigger DESC, evolved.national_order ASC").
		Select()
	if err != nil {
		return err
	}

	for _, p := range pokemon {
		p.LoadDexNumbers(dexNumbers)
		p.LoadEvolutions(evolutions)
	}

	return c.JSON(http.StatusOK, pokemon)
}

func retrieveHandler(c echo.Context) error {
	app := c.Get("app").(*application.App)

	id := c.Param("id")

	var p Pokemon
	var dexNumbers []*dexnumbers.GameFamilyDexNumber
	var evolutions []*Evolution

	err := app.DB.Model(&p).Relation("GameFamily").Where("pokemon.id = ?", id).First()
	if err != nil {
		return err
	}
	err = app.DB.Model(&dexNumbers).Where("pokemon_id = ?", id).Select()
	if err != nil {
		return err
	}
	err = app.DB.
		Model(&evolutions).
		Relation("EvolvingPokemon").
		Relation("EvolvedPokemon").
		Join("INNER JOIN pokemon AS evolved ON evolutions.evolved_pokemon_id = evolved.id").
		Join("INNER JOIN pokemon AS evolving ON evolutions.evolving_pokemon_id = evolving.id").
		Where("evolutions.evolution_family_id = ?", p.EvolutionFamilyID).
		OrderExpr("CASE WHEN trigger = 'breed' THEN evolving.national_id ELSE evolved.national_id END, trigger DESC, evolved.national_order ASC").
		Select()
	if err != nil {
		return err
	}

	p.LoadDexNumbers(dexNumbers)
	p.LoadEvolutions(evolutions)

	return c.JSON(http.StatusOK, p)
}
