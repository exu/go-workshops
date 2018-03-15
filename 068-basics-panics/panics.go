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

// program is only valid when there is NAME env set
func main() {
	println("Starting", name)
}
