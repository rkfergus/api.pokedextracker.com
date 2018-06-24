package application

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Middleware(app *App) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := app.WithContext(r.Context())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
