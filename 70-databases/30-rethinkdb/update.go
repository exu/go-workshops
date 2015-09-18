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
type ScoreEntry struct {
	ID         string `gorethink:"id,omitempty"`
	PlayerName string
	Score      int
}

func main() {
	fmt.Println("Connecting to RethinkDB")

	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "players",
	})

	session.SetMaxOpenConns(10)

	if err != nil {
		log.Fatal("Could not connect")
	}

	fmt.Println("Loop through updates")
	for {
		var scoreentry ScoreEntry
		pl := rand.Intn(10000)
		sc := rand.Intn(6) - 2
		res, err := r.Table("scores").Get(strconv.Itoa(pl)).Run(session)
		if err != nil {
			log.Fatal(err)
		}

		err = res.One(&scoreentry)
		scoreentry.Score = scoreentry.Score + sc

		startingTime := time.Now()
		_, err = r.Table("scores").Update(scoreentry).RunWrite(session)
		fmt.Println(time.Now().Sub(startingTime))

		fmt.Println("Updating ", scoreentry)
	}
}
