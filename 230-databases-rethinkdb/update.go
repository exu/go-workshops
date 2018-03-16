package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	r "gopkg.in/dancannon/gorethink.v1"
)

//ScoreEntry for scores
// type ScoreEntry struct {
// 	ID         string `gorethink:"id,omitempty"`
// 	PlayerName string
// 	Score      int
// }

func main() {
	fmt.Println("Connecting to RethinkDB")

	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:7704",
		Database: "players",
	})

	session.SetMaxOpenConns(10)

	if err != nil {
		log.Fatal("Could not connect")
	}

	fmt.Println("Loop through updates")
	for {
		pl := rand.Intn(1000)
		sc := rand.Intn(6) - 2

		// startingTime := time.Now()
		_, err = r.Table("scores").Get(strconv.Itoa(pl)).Update(map[string]interface{}{
			"Score": r.Row.Field("Score").Add(sc),
		}).RunWrite(session)
		// fmt.Println(time.Now().Sub(startingTime))

		time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
	}
}
