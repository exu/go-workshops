package main

import (
	"flag"
	"fmt"
)

// Przykład 2: Two flags deals same value,
// be sure that they are set in `init()` function
var size string

func init() {
	const (
		defaultSize = "1M"
		usage       = "Size in Megabytes"
	)
	flag.StringVar(&size, "size", defaultSize, usage)
	flag.StringVar(&size, "s", defaultSize, usage+" - little shorter :)")
	flag.Parse()
}

func main() {
	fmt.Println("Rozmiar:", size)
}
