package routes

import (
	"io/ioutil"
	"net/http"
	"scanbu-api/modules/facebook-bot/handlers"

	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/pressly/chi"
)

func fbotRoute(route *chi.Mux) {
	route.Get("/fbot", handlers.FacebookBot)
	route.Post("/fbot", handlers.FacebookBotReceiver)

	route.Get("/fb/politica/privacidade", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadFile("politica-de-privacidade.txt")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "%s", b)
	})
}
