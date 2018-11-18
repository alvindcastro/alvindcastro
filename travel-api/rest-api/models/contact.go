package models

import "gopkg.in/mgo.v2/bson"

type Contact struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Prenom    string        `bson:"prenom" json:"prenom"`
	Nom       string        `bson:"nom" json:"nom"`
	Telephone string        `bson:"telephone" json:"telephone"`
}
