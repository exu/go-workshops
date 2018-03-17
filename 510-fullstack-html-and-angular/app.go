package main

import (
	"github.com/gin-gonic/gin"
)

type Point struct {
	X int
	Y int
}

type Hash map[string]interface{}

type Map struct {
	Id         int
	Name       string
	Tags       []string
	KeyPoints  []Point
	Properties Hash
}

var response Map

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
	response.Id++
	c.HTML(200, "index.html", response)
}

func main() {
	app := gin.Default()
	// add templates
	app.LoadHTMLFiles("./static/index.html")
	// add static handler (could be done by some http server i.e. nginx)
	app.Static("/static", "static")
	app.GET("/", handler)
	app.Run(":8080")
}
