package pokemon

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pokedextracker/api.pokedextracker.com/application"
	"github.com/pokedextracker/api.pokedextracker.com/dexnumbers"
)

func listHandler(w http.ResponseWriter, r *http.Request) {
	app := application.FromContext(r.Context())

	var pokemon []*Pokemon
	var dexNumbers []*dexnumbers.GameFamilyDexNumber
	var evolutions []*Evolution

	app.DB.Order("national_order ASC").Find(&pokemon)
	app.DB.Find(&dexNumbers)
	app.DB.
		Joins("INNER JOIN pokemon AS evolved ON evolutions.evolved_pokemon_id = evolved.id").
		Joins("INNER JOIN pokemon AS evolving ON evolutions.evolving_pokemon_id = evolving.id").
		Order("CASE WHEN trigger = 'breed' THEN evolving.national_id ELSE evolved.national_id END, trigger DESC, evolved.national_order ASC").
		Find(&evolutions)

	for _, p := range pokemon {
		p.LoadDexNumbers(dexNumbers)
		p.LoadEvolutions(evolutions)
	}

	json.NewEncoder(w).Encode(pokemon)
}

func retrieveHandler(w http.ResponseWriter, r *http.Request) {
	app := application.FromContext(r.Context())

	id := mux.Vars(r)["id"]

	var p Pokemon
	var dexNumbers []*dexnumbers.GameFamilyDexNumber
	var evolutions []*Evolution

	app.DB.Where("id = ?", id).First(&p)
	app.DB.Where("pokemon_id = ?", id).Find(&dexNumbers)
	app.DB.
		Joins("INNER JOIN pokemon AS evolved ON evolutions.evolved_pokemon_id = evolved.id").
		Joins("INNER JOIN pokemon AS evolving ON evolutions.evolving_pokemon_id = evolving.id").
		Where("evolutions.evolution_family_id = ?", p.EvolutionFamilyID).
		Order("CASE WHEN trigger = 'breed' THEN evolving.national_id ELSE evolved.national_id END, trigger DESC, evolved.national_order ASC").
		Find(&evolutions)

	p.LoadDexNumbers(dexNumbers)
	p.LoadEvolutions(evolutions)

	json.NewEncoder(w).Encode(p)
}
