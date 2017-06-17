package routes

import (
	"scanbu-api/modules/facebook-bot/handlers"

	"github.com/pressly/chi"
)

func fbotRoute(route *chi.Mux) {
	route.Get("/fbot", handlers.FacebookBot)
}
