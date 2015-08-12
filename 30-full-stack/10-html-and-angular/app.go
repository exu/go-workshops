package main

import (
	// wykorzystamy gin prosty router z kilkoma middleware'ami
	"github.com/gin-gonic/gin"
)

// strukturka
type Point struct {
	X int
	Y int
}

// zdefiniujmy sobie hash
// szkoda że ten typ danych nie
// jest wbudowany
type Hash map[string]interface{}

// i jeszcze jedna
type Map struct {
	Id         int
	Name       string
	Tags       []string
	KeyPoints  []Point
	Properties Hash
}

var response Map

// init wykonywany jest przy pierwszym zaczytaniu paczki
func init() {
	response = Map{
		Id:        1000,
		Name:      "super response",
		Tags:      []string{"Góry", "Doliny"},
		KeyPoints: []Point{{1, 2}, {2, 3}},
		Properties: Hash{
			"poziom_trudności": "super skomplikowany",
			"typ_nawierzchni":  "głównie asfalt",
		},
	}
}

func handler(c *gin.Context) {
	// zwiększamy wartość
	// to jest app server więc między requestami
	// się zwiększy również
	response.Id++

	// render index.html -> patrz LoadHTMLGlob w main
	c.HTML(200, "index.html", response)
}

func main() {
	app := gin.Default()

	// gdzie ma szukać templatek (tu w tym samym katalogu
	// w którym odpalamy program
	app.LoadHTMLGlob("*.html")

	// handler static, ja osobiście wolę
	// użyć do tego nginxa (tutaj uproszczenie na
	// potrzeby prezentacji)
	app.Static("/static", "static")

	app.GET("/", handler)

	app.Run(":8080")
}
