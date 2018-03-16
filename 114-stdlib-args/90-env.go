package main

import (
	"os"
)

var port = os.Getenv("PORT")

func main() {
	println("Starting", port)
}
