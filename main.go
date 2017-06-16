package main

import (
	"net/http"
	"scanbu-api/helpers"
	"scanbu-api/helpers/database"
	"scanbu-api/modules/data-extractor/lib"
	"scanbu-api/routes"

	log "github.com/Sirupsen/logrus"
)

func main() {
	port := helpers.GetENVorDefault("PORT", "8000")

	database.StartDB()
	defer database.CloseDB()

	go lib.ExtractorProcess()

	log.Info("Server started at: ", port)

	log.Fatal(http.ListenAndServe(":"+port, routes.R))
}
