package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Lets define some loggers
var (
	Trace   *log.Logger // details
	Info    *log.Logger // important informations
	Warning *log.Logger // warnings
	Error   *log.Logger // fuckups
)

func main() {
	// warnings setup
	warningsFile, err := os.OpenFile("warnings.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open warning log file")
	}
	defer warningsFile.Close()

	// errors setup
	errorsFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open errors log file")
	}
	defer errorsFile.Close()

	// multi writer will log to a fail and to stdout in parallel.
	multiWriter := io.MultiWriter(errorsFile, os.Stderr)

	Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ltime)

	Warning = log.New(warningsFile,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(multiWriter,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Trace.Println("I have something standard to say.")
	Info.Println("Important Information.")
	Warning.Println("There is something you need to know about.")
	Error.Println("Something has failed.")
}
