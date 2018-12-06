package model

import (
	"gopkg.in/mgo.v2/bson"
)

// Shopping - Model
type Shopping struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	User       int           `bson:"user" json:"user"`
	Products   []string      `bson:"products" json:"products"`
	Payment    string        `bson:"payment" json:"payment"`
	PriceTotal int           `bson:"price_total" json:"price_total"`
}

// ShoppingID - for request
type ShoppingID struct {
	ID string `json:"id"`
}
