package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Item struct {
	Id          bson.ObjectId `bson:"_id"`
	Name        string
	Description string
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("goworkshops_tailed").C("items")

	iter := c.Find(nil).Sort("$natural").Tail(12 * time.Second)
	defer iter.Close()

	var result Item
	var lastId bson.ObjectId

	for {
		for iter.Next(&result) {
			fmt.Printf("%+v\n\n", result)
			lastId = result.Id
		}
		if iter.Err() != nil {
			iter.Close()
			fmt.Println(iter.Err().Error())
			time.Sleep(time.Second)
		}
		if iter.Timeout() {
			continue
		}
		query := c.Find(bson.M{"_id": bson.M{"$gt": lastId}})
		iter = query.Sort("$natural").Tail(5 * time.Second)
	}
}
