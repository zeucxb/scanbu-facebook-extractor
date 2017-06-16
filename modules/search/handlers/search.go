package handlers

import (
	"fmt"
	"net/http"
	"scanbu-api/helpers"
	"scanbu-api/modules/product/models"
	"strings"

	"gopkg.in/mgo.v2/bson"

	log "github.com/Sirupsen/logrus"
)

// Search is the search handler
func Search(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	if keywords, ok := q["keyword"]; ok {
		search := strings.Join(keywords, " ")

		var products []models.Product
		query := bson.M{
			"message": bson.M{
				"$regex": bson.RegEx{
					Pattern: fmt.Sprintf("\b%s\b", search),
					Options: "i",
				},
			},
		}

		if err := models.Products().Find(query).All(&products); err == nil {
			if len(products) > 0 {
				bytes, err := helpers.JSONMarshal(products, true)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Fprintf(w, "%s", bytes)
			} else {
				query = bson.M{
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
						log.Fatal(err)
					}

					fmt.Fprintf(w, "%s", bytes)
				}
			}
		}
	}
}
