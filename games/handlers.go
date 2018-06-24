package games

import (
	"encoding/json"
	"net/http"

	"github.com/pokedextracker/api.pokedextracker.com/application"
)

func listHandler(w http.ResponseWriter, r *http.Request) {
	app := application.FromContext(r.Context())

	var games []Game

	app.DB.
		Joins("INNER JOIN game_families ON games.game_family_id = game_families.id").
		Where("game_families.published = TRUE").
		Order("game_families.order DESC, games.order ASC").
		Find(&games)

	json.NewEncoder(w).Encode(games)
}
