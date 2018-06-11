package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func collection() *mongo.Collection {
	client, err := mongo.NewClient("mongodb://localhost:7702")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("official").Collection("qux")
}
func main() {
	collection := collection()
	res, err := collection.InsertOne(context.Background(), map[string]string{"hello": "world"})
	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID

	fmt.Printf("%+v\n", id)
}
