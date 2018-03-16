package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:7703",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func main() {
	err := client.Set("sessions", 10, 0).Err()
	if err != nil {
		panic(err.Error())
	}

	sessions, _ := client.Get("sessions").Int64()

	fmt.Println("Active sessions:", sessions)
}
