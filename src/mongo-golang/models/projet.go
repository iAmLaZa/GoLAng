package models

import "gopkg.in/mgo.v2/bson"

type Projet struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Link        string        `json:"link" bson:"link"`
	Description string        `json:"description" bson:"description"`
}
