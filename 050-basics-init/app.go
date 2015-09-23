package main

import (
	"fmt"
	"github.com/exu/go-workshops/050-basics-init/screamer"
	"os"
)

func main() {
	fmt.Println("Start")
	screamer.ScreamAAAA(os.Stdout)
}
