package main

import (
	"flag"
	"fmt"
)

// Przykład 1: Pojedyńcza flaga typu string
// Referencja!
var logLevel = flag.String("log_level", "warning", "Ustawia poziom logowania")

func main() {
	// musimy wywołać flag parse
	flag.Parse()
	fmt.Println("Poziom logowania:", *logLevel)
}
