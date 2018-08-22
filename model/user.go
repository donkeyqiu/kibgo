package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		ID       bson.ObjectId `bson:"_id,omitempty"`
		Username string        `bson:"username"`
		Userid   string        `bson:"userid"`
		Pic      string        `bson:"pic"`
		Email    string        `bson:"email"`
		Mobile   string        `bson:"mobile"`
		Password string        `bson:"password"`
		Sex      bool          `bson:"sex"`
		Status   bool          `bson:"status"`
		Salt     string        `bson:"-"`
	}
)
