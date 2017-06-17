package routes

import (
	"scanbu-api/modules/facebook-bot/handlers"

	"github.com/pressly/chi"
)

<<<<<<< HEAD
func fbotRoute(route *chi.Mux) {
=======
func searchRoute(route *chi.Mux) {
>>>>>>> 1f1ef17b19fc79941108eaf7cf2fe321a479b32f
	route.Get("/fbot", handlers.FacebookBot)
}
