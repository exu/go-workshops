package main

import (
	"fmt"
	"log"

	r "gopkg.in/dancannon/gorethink.v1"
)

func main() {
	fmt.Println("I'm watching You!")

	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:7704",
		Database: "players",
	})

	if err != nil {
		log.Fatal(err)
	}

	res, err := r.Table("scores").Changes().Run(session)

	if err != nil {
		log.Fatalln(err)
	}

	var value interface{}
	for res.Next(&value) {
		fmt.Println(value)
	}
}
