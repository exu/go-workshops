package main

import (
	"flag"
	"fmt"
)

// Przykład 1: single flag of string type
// Referencja!
var logLevel = flag.String("log_level", "warning", "Set logging level")

func main() {
	// we need to call Parse(), after that flags will be parsed
	flag.Parse()
	fmt.Println("Current log level:", *logLevel)
}
