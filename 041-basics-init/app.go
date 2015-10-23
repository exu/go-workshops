package main

import (
	"fmt"
	"github.com/exu/go-workshops/041-basics-init/package1"
	"github.com/exu/go-workshops/041-basics-init/package2"
	"github.com/exu/go-workshops/041-basics-init/package3"
	"os"
)

func main() {
	fmt.Println("Start")
	package1.Write(os.Stdout)
	package2.Write(os.Stdout)

	fmt.Println("Writing all inside package3")
	package3.Write(os.Stdout)
}
