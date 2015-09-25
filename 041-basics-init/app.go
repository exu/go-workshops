package main

import (
	"fmt"
	"github.com/exu/go-workshops/041-basics-init/screamer"
	"github.com/exu/go-workshops/041-basics-init/silencer"
	"os"
)

func main() {
	fmt.Println("Start")
	screamer.ScreamAAAA(os.Stdout)
	silencer.Whisp(os.Stdout)
}
