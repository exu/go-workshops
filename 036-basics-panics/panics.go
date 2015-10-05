package main

import (
	"fmt"
	"os"
)

var name = os.Getenv("NAME")

func init() {
	if name == "" {
		panic("no value for $NAME")
	}

}

func main() {
	fmt.Printf("Starting %s", name)
}
