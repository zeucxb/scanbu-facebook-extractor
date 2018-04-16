package routes

import (
	"scanbu-extractor/modules/search/handlers"

	"github.com/pressly/chi"
)

func searchRoute(route *chi.Mux) {
	route.Get("/search", handlers.Search)
}
