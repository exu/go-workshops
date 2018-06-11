package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("localhost:7702")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("goworkshops_fetch").C("people")
	c.DropCollection()

	err = c.Insert(
		Person{"Ale", "+55 53 8116 9639"},
		Person{"Cla", "+55 53 8402 8510"},
	)

	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
