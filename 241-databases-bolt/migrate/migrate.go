package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"gopkg.in/redis.v3"
)

var (
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func main() {
	keys, err := client.Keys("*").Result()
	if err != nil {
		log.Fatal(err, keys)
	}

	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key)
	}

	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("MyBucket"))
		fmt.Printf("%+v", b)

		return b.Put([]byte("answer"), []byte("42"))
	})
}
