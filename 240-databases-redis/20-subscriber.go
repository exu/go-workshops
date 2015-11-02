package main

import (
	"gopkg.in/redis.v3"
	"log"
)

var (
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:7703",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
}

func main() {

	pubsub, err := client.Subscribe("mychannel")
	if err != nil {
		panic(err)
	}
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			panic(err)
		}

		log.Println(msg.Channel, msg.Payload)
	}

}
