package main

import (
	"net/http"
	"scanbu-extractor/helpers"
	"scanbu-extractor/helpers/database"
	"scanbu-extractor/modules/data-extractor/lib"
	"scanbu-extractor/routes"

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
