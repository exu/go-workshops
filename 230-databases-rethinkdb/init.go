package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

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
		Address:  "localhost:7704",
		Database: "players",
	})

	session.SetMaxOpenConns(10)

	if err != nil {
		log.Fatal("Could not connect")
	}

	r.DBDrop("players").Exec(session)
	err = r.DBCreate("players").Exec(session)
	if err != nil {
		panic(err)
	}

	err = r.DB("players").TableDrop("scores").Exec(session)
	err = r.DB("players").TableCreate("scores").Exec(session)
	if err != nil {
		log.Fatal("Could not create table", err)
	}

	err = r.DB("players").Table("scores").IndexCreate("Score").Exec(session)
	if err != nil {
		log.Fatal("Could not create index")
	}

	fmt.Println("Inserting new records, quite slow need to be batch :| ")
	for i := 0; i < 1000; i++ {
		player := new(ScoreEntry)
		player.ID = strconv.Itoa(i)
		player.PlayerName = fmt.Sprintf("Player %d", i)
		player.Score = rand.Intn(100)
		_, err := r.Table("scores").Insert(player).RunWrite(session)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Inserted player number: ", i)
	}
}
