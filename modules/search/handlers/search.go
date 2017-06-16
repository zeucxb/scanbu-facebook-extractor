package handlers

import (
	"fmt"
	"net/http"
	"scanbu-api/helpers"
	"scanbu-api/modules/data-extractor/lib"
	"scanbu-api/modules/product/models"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

var groups = []string{
	"193939064109587",
	"1088976661131866",
	"415451778499368",
}

// Search is the search handler
func Search(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	if keywords, ok := q["keyword"]; ok {
		search := strings.Join(keywords, " ")

		var products []models.Product
		query := bson.M{
			"message": bson.M{
				"$regex": bson.RegEx{
					Pattern: search,
					Options: "i",
				},
			},
		}

		if err := models.Products().Find(query).All(&products); err == nil {
			bytes, err := helpers.JSONMarshal(products, true)
			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, "%s", bytes)
		}

		go lib.Proccess(groups)
	}
}