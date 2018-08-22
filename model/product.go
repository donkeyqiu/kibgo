package model

import "gopkg.in/mgo.v2/bson"

type (
	Product struct {
		ID           bson.ObjectId `bson:"_id,omitempty"`
		Barcode      string        `bson:"barcode"`
		Pic          string        `bson:"pic"`
		Price        int           `bson:"price"`
		Cost         float32       `bson:"cost"`
		TagPrice     int           `bson:"tag_price"`
		SpecialPrice int           `bson:"special_price"`
		Num          int           `bson:"num"`
	}
)
