package lib

import (
	"fmt"
	"scanbu-extractor/modules/product/models"

	"gopkg.in/mgo.v2/bson"
)

// Search the product and returns
func Search(search string) (products []models.Product, err error) {
	query := bson.M{
		"message": bson.M{
			"$regex": bson.RegEx{
				Pattern: fmt.Sprintf("\b%s\b", search),
				Options: "i",
			},
		},
	}

	if err = models.Products().Find(query).Sort("-createdTime").All(&products); err == nil {
		if len(products) > 0 {
			return products, nil
		}

		query = bson.M{
			"message": bson.M{
				"$regex": bson.RegEx{
					Pattern: search,
					Options: "i",
				},
			},
		}

		if err = models.Products().Find(query).Sort("-createdTime").All(&products); err == nil {
			return products, nil
		}

		return
	}

	return
}
