package dispatcher

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Dispatch(DB *sql.DB) http.Handler {
	r := chi.NewRouter()

	//Logging every request
	r.Use(middleware.Logger)

	r.Route("/app",func(r chi.Router) {
		r.Mount()
	})



}