package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

func main() {
	// db.createCollection("messages", { size: 100000000, capped: true })
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("test").C("cars")
}
