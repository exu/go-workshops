package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

func main() {
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

	iterations := 1000

	now := time.Now()
	for i := 0; i < iterations; i++ {
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("MyBucket"))
			key := fmt.Sprintf("answer_%d", i)
			return b.Put([]byte(key), []byte(strconv.Itoa(rand.Int())))
		})
	}
	fmt.Println(time.Since(now))

	now = time.Now()

	for i := 0; i < iterations; i++ {
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("MyBucket"))
			key := fmt.Sprintf("answer_%d", i)
			v := b.Get([]byte(key))
			fmt.Printf("The answer is: %s\n", v)
			return nil
		})
	}

	fmt.Println(time.Since(now))
}
