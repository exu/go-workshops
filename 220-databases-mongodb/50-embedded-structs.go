package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	COLOR_WHITE = iota
	COLOR_RED
	COLOR_GREEN
	COLOR_BLUE
	COLOR_BLACK
)

type Person struct {
	Name        string
	Phone       string
	Description string
}

type Door struct {
	HasMirros  bool
	HasHandler bool
	HasGlass   bool
}

type Car struct {
	Id         bson.ObjectId `bson:"_id"`
	Make       string
	Model      string
	Color      int
	Doors      []Door
	Passengers []Person
	SeatsCount int
}

func main() {
	session, err := mgo.Dial("localhost:7702")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("test").C("cars")
	c.DropCollection()

	kazik := Person{"Kazik", "+4890909090", "PHP Deweloper"}
	stasiu := Person{"Stasiu", "+4888889999", "Taki tam"}

	id := bson.NewObjectId()
	mustang := &Car{
		Id:    id,
		Make:  "Ford",
		Model: "Mustang GT500",
		Color: COLOR_BLACK,
		Doors: []Door{
			{true, true, true},
			{true, true, true},
			{true, false, false},
			{true, false, false},
		},
		Passengers: []Person{stasiu, kazik},
		SeatsCount: 10,
	}

	// put car to database
	err = c.Insert(&mustang)
	if err != nil {
		log.Fatal(err)
	}

	// now we can get our fresh record from DB
	var result Car
	c.FindId(id).One(&result)
	fmt.Printf("%+v", result)
}
