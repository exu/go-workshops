package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Item struct {
	Id   int `bson:"_id"`
	Name string
	Tags []mgo.DBRef
}

type Tag struct {
	Id   int `bson:"_id"`
	Name string
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	db1 := session.DB("db1")
	tags := db1.C("tags")
	items := db1.C("items")

	tags.Insert(bson.M{"_id": 1, "name": "super"})
	tags.Insert(bson.M{"_id": 2, "name": "słaby"})
	tags.Insert(bson.M{"_id": 3, "name": "in the middle"})

	ref1 := mgo.DBRef{Collection: "tags", Id: 1}
	ref2 := mgo.DBRef{Collection: "tags", Id: 2}

	items.Insert(Item{Id: 1, Tags: []mgo.DBRef{ref1, ref2}})
	items.Insert(Item{Id: 2, Tags: []mgo.DBRef{ref2}})

	result := Tag{}

	// tag id=1
	err = db1.FindRef(&ref1).One(&result)
	fmt.Println(result)

	// get reference
	item := Item{}
	err = items.Find(bson.M{"_id": 1}).One(&item)
	if err != nil {
		fmt.Println(err)
	}

	err = db1.FindRef(&item.Tags[1]).One(&result)

	// tag id=2
	fmt.Println(result)
}
