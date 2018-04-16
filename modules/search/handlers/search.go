package handlers

import (
	"fmt"
	"net/http"
	"scanbu-extractor/helpers"
	"scanbu-extractor/modules/search/lib"

	log "github.com/Sirupsen/logrus"
)

// Search is the search handler
func Search(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	if keywords, ok := q["keyword"]; ok {
		search := keywords[0]

		products, err := lib.Search(search)
		if err != nil {
			log.Fatal(err)
		}

		bytes, err := helpers.JSONMarshal(products, true)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "%s", bytes)
	}
}
