package main

import (
	"log"

	"github.com/go-redis/redis"
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

	pubsub := client.Subscribe("mychannel")
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			panic(err)
		}

		log.Println(msg.Channel, msg.Payload)
	}

}
