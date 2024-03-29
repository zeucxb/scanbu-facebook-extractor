package routes

import (
	"github.com/goware/cors"
	"github.com/pressly/chi"
)

// R is the router object
var R = chi.NewRouter()

func init() {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	R.Use(cors.Handler)

	searchRoute(R)
	fbotRoute(R)
}
