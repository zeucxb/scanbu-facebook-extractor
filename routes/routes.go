package routes

import "github.com/pressly/chi"

// R is the router object
var R = chi.NewRouter()

func init() {
	searchRoute(R)
}
