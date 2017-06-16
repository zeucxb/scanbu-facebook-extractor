package main

import (
	"net/http"
	"scanbu-api/helpers"
	"scanbu-api/helpers/database"
	"scanbu-api/routes"
)

func main() {
	port := helpers.GetENVorDefault("PORT", "8000")

	database.StartDB()
	defer database.CloseDB()

	http.ListenAndServe(":"+port, routes.R)
}
