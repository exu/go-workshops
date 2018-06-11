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

type Hello struct {
	Hello string `bson:"hello"`
}

func main() {

	collection := collection()
	cur, err := collection.Find(context.Background(), map[string]string{"hello": "world"})

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	hell := Hello{}

	for cur.Next(context.Background()) {
		err := cur.Decode(&hell)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", hell)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
