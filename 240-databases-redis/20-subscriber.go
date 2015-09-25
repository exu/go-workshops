package main

import (
	"fmt"
	"gopkg.in/redis.v3"
)

var (
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
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

		fmt.Println(msg.Channel, msg.Payload)
	}

}
