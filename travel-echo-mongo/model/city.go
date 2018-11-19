package model

import "gopkg.in/mgo.v2/bson"

type (
	City struct {
		ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name        string        `json:"name" bson:"name"`
		Desc        string        `json:"desc" bson:"desc"`
		Attractions []string      `json:"attractions,omitempty" bson:"attractions,omitempty"`
	}
)
