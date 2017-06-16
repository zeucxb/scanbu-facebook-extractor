package routes

import (
	"scanbu-api/modules/search/handlers"

	"github.com/pressly/chi"
)

func searchRoute(route *chi.Mux) {
	route.Get("/search", handlers.Search)
}
