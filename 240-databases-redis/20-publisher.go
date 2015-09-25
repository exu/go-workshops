package main

import (
	"gopkg.in/redis.v3"
)

var (
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	err := client.Publish("mychannel", "hello").Err()
	if err != nil {
		panic(err)
	}
}
