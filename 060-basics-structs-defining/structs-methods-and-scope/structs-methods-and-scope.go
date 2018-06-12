package main

import (
	"fmt"

	"github.com/exu/go-workshops/060-basics-structs-defining/image"
)

func main() {
	img := image.NewImage(1888)

	fmt.Printf("%+v", img.GetSize())
}
