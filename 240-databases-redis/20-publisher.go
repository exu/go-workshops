package main

import (
	"os"
	"time"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:7703",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	if len(os.Args) < 2 {
		println("Podaj wiadomość")
		return
	}

	for {
		err := client.Publish("mychannel", "hello "+os.Args[1]).Err()
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond * 500)
	}
}
