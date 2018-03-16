package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Item struct {
	Id          bson.ObjectId `bson:"_id"`
	Name        string
	Description string
}

func main() {
	// db.createCollection("items", { size: 1000, capped: true })
	session, err := mgo.Dial("localhost:7702")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	db := session.DB("goworkshops_tailed")
	c := db.C("items")
	info := &mgo.CollectionInfo{
		Capped:   true,
		MaxDocs:  1000,
		MaxBytes: 1024 * 1024 * 100,
	}

	c.DropCollection()
	c.Create(info)

	c.Insert(Item{bson.NewObjectId(), "Kij", "Samobij"})
	c.Insert(Item{bson.NewObjectId(), "Okulary", "Ciemne, czerwone"})
	c.Insert(Item{bson.NewObjectId(), "Pączek", "ugryziony z lewej strony"})
}

// should give you
// > db.items.stats().capped
// true
// in mongo shell
