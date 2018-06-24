package games

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router) {
	s := r.PathPrefix("/games").Subrouter()

	s.HandleFunc("", listHandler)
}
