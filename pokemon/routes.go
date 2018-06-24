package pokemon

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router) {
	s := r.PathPrefix("/pokemon").Subrouter()

	s.HandleFunc("", listHandler)
	s.HandleFunc("/{id}", retrieveHandler)
}
