package service

import (
	"gopkg.in/mgo.v2"

	"log"
)

type BaseModel struct {
	Server   string
	Database string
}

type Msg map[string]interface{}

var db *mgo.Database

// Establish a connection to database
func (m *BaseModel) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}