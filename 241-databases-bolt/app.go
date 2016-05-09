package main

import (
	"log"

	"fmt"
	"github.com/boltdb/bolt"
	"math/rand"
	"strconv"
	"time"
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

	now := time.Now()
	for i := 0; i < 1000; i++ {
		db.Update(func(tx *bolt.Tx) error {

			b := tx.Bucket([]byte("MyBucket"))
			// fmt.Printf("%+v\n", b)
			val := rand.Int()
			return b.Put([]byte("answer"), []byte(strconv.Itoa(val)))
		})
	}
	fmt.Println(time.Since(now))

	now = time.Now()

	for i := 0; i < 1000000; i++ {
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("MyBucket"))
			b.Get([]byte("answer"))
			//			fmt.Printf("The answer is: %s\n", v)
			return nil
		})
	}

	fmt.Println(time.Since(now))
}
