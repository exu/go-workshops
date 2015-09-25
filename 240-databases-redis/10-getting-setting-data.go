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
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func main() {
	key := "sessions"

	err := client.Set(key, 10, 0).Err()
	if err != nil {
		panic(err.Error())
	}

	sessions, _ := client.Get(key).Int64()

	fmt.Println("Active sessions:", sessions)
}
