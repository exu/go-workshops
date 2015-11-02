package main

import (
	"os"
)

var name = os.Getenv("NAME")

func init() {
	if name == "" {
		panic("no value for $NAME")
	}

}

func main() {
	println("Starting", name)
}
