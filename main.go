package main

import (
	"net/http"
	"log"

	."./config"
	."./service"
	."./router"
	."./controller"

	"gopkg.in/mgo.v2"
	)

var config = Config{}
var base = BaseModel{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	go HandleMessages()
	config.Read()
	base.Server = config.Server
	base.Database = config.Database
	base.Connect()
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}