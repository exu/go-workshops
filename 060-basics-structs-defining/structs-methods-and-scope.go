package main

import (
	"./image"
	"fmt"
)

func main() {
	img := image.NewImage(1888)

	fmt.Printf("%+v", img.GetSize())
}
