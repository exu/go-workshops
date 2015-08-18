package main

import (
	"flag"
	"fmt"
)

// Przykład 2: Dwie flagi dzielą jedną wartość
// Upewnij się że obie mają tego samego defaulta
// Muszą być ustawione na poziomie funkcji init()
var size string

func init() {
	const (
		defaultSize = "mini"
		usage       = "Rozmiar w Megabajtach"
	)
	flag.StringVar(&size, "size", defaultSize, usage)
	flag.StringVar(&size, "s", defaultSize, usage+" (troszkę krócej)")
}

func main() {
	flag.Parse()
	fmt.Println("Rozmiar:", size)
}
