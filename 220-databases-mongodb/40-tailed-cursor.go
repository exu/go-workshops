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
	session, err := mgo.Dial("localhost:7702")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("goworkshops_tailed").C("items")

	var result Item

	lastId := bson.NewObjectId()
	for {
		println("starting")
		query := c.Find(bson.M{"_id": bson.M{"$gt": lastId}})
		iter := query.Sort("$natural").Tail(5 * time.Second)

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
			println("timeout")
			continue
		}
	}
}
