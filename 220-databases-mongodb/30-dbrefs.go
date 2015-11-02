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
	session, err := mgo.Dial("localhost:7702")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	const DB1NAME = "goworkshops_ref_db1"
	const DB2NAME = "goworkshops_ref_db2"

	db1 := session.DB(DB1NAME)
	db2 := session.DB(DB2NAME)

	tags1 := db1.C("tags")
	tags1.DropCollection()
	tags1.Insert(bson.M{"_id": 1, "name": "super"})
	tags1.Insert(bson.M{"_id": 2, "name": "słaby"})
	tags1.Insert(bson.M{"_id": 3, "name": "in the middle"})

	tags2 := db2.C("tags")
	tags2.DropCollection()
	tags2.Insert(bson.M{"_id": 1, "name": "zielony"})
	tags2.Insert(bson.M{"_id": 2, "name": "czerwony"})
	tags2.Insert(bson.M{"_id": 3, "name": "cytrynowa malaga"})

	ref1 := mgo.DBRef{Database: DB1NAME, Collection: "tags", Id: 1}
	ref2 := mgo.DBRef{Database: DB2NAME, Collection: "tags", Id: 2}

	items := db1.C("items")
	items.DropCollection()
	items.Insert(Item{Id: 1, Tags: []mgo.DBRef{ref1, ref2}})
	items.Insert(Item{Id: 2, Tags: []mgo.DBRef{ref2}})

	result := Tag{}

	// tag id=1
	err = session.FindRef(&ref1).One(&result)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("ref1 result:", result)

	err = session.FindRef(&ref2).One(&result)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("ref2 result:", result)

	// get reference
	fmt.Println("Find ref from db:")
	item := Item{}
	err = items.Find(bson.M{"_id": 1}).One(&item)
	if err != nil {
		fmt.Println(err)
	}

	err = session.FindRef(&item.Tags[1]).One(&result)

	// tag id=2
	fmt.Println("item.Tags[1]", result)
}
