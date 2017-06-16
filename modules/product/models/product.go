package models

import (
	"scanbu-api/helpers/database"

	mgo "gopkg.in/mgo.v2"
)

// Products returns the product collection
func Products() *mgo.Collection {
	return database.Collection("products", []mgo.Index{
		mgo.Index{
			Key: []string{"name", "title", "description", "price", "place", "createdTime"},
		},
		mgo.Index{
			Key:    []string{"facebookId"},
			Unique: true,
		},
	})
}

// Product is the product type
type Product struct {
	ID          string      `json:"_id" bson:"_id"`
	Message     string      `json:"message" bson:"message" facebook:"message"`
	Title       string      `json:"title" bson:"title"`
	Description string      `json:"description" bson:"description"`
	Prince      float64     `json:"price" bson:"price"`
	Place       string      `json:"place" bson:"place"`
	Type        string      `json:"type" bson:"type" facebook:"type"`
	Link        string      `json:"link" bson:"link" facebook:"permalink_url"`
	Picture     string      `json:"picture" bson:"picture" facebook:"picture"`
	FullPicture string      `json:"fullPicture" bson:"fullPicture" facebook:"full_picture"`
	CreatedTime string      `json:"createdTime" bson:"createdTime" facebook:"created_time"`
	FacebookID  string      `json:"facebookId" bson:"facebookId" facebook:"id"`
	User        ProductUser `json:"user" bson:"user" facebook:"from"`
	Group       Group       `json:"group" bson:"group" facebook:"target"`
	Attachments Attachments `json:"attachments" bson:"attachments" facebook:"attachments.data"`
}

// ProductUser is the user that make the post
type ProductUser struct {
	Name       string `json:"name" bson:"name" facebook:"name"`
	FacebookID string `json:"facebookId" bson:"facebookId" facebook:"id"`
}

// Group is the group where the product was posted
type Group struct {
	Name       string `json:"name" bson:"name" facebook:"name"`
	Pricacy    string `json:"privacy" bson:"privacy" facebook:"privacy"`
	FacebookID string `json:"facebookId" bson:"facebookId" facebook:"id"`
}

// Attachments are the product attachments
type Attachments struct {
	Media []Media `json:"media" bson:"media" facebook:"media"`
}

// Media is a product media
type Media struct {
	Image Image `json:"image" bson:"image" facebook:"image"`
}

// Image is the media image
type Image struct {
	Width  string `json:"width" bson:"width" facebook:"width"`
	Height string `json:"height" bson:"height" facebook:"height"`
	Src    string `json:"src" bson:"src" facebook:"src"`
}
