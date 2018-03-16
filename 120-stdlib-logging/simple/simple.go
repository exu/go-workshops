package main

import (
	"log"
	"os"
)

// we can init our logger with some
// more sophisticated config
func init() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("TRACE: ")
	/*
	   Ldate		   // the date: 2009/01/23
	   Ltime           // the time: 01:23:23
	   Lmicroseconds   // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	   Llongfile       // full file name and line number: /a/b/c/d.go:23
	   Lshortfile      // final file name element and line number: d.go:23. overrides Llongfile
	   LstdFlags       // Ldate | Ltime // initial values for the standard logger
	*/
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Println("Starting my super app! Please fasten your seatbelts")

	names := []string{"Tytus", "Romek", "Atomek"}
	log.Printf("Loading users %+v\n", names)

	log.Fatalln("Fatality!")
	// calls `os.Exit(1)` => exit status 1
	log.Println("Ending app")
	// never will be called
}
