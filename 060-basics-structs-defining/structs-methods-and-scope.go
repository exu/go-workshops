package main

import (
	"./image"
	"fmt"
)

func main() {
	img := image.Image{
		URL: "http://www.wp.pl/logo.png",
		// contentType: "image/png",
		// size: 12999
	}

	img.SetSize(10000)
	fmt.Printf("%+v", img.GetSize())
}
