package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Nasze loggery
var (
	Trace   *log.Logger // detale
	Info    *log.Logger // ważne informacje
	Warning *log.Logger // ostrzeżenia
	Error   *log.Logger // wyjebki
)

// initLog inicjuje nasze loggery
func initLog(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// pliczek dla warningów
	warnings, err := os.OpenFile("warnings.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open warning log file")
	}
	defer warnings.Close()

	// pliczek dla errorów
	errors, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open errors log file")
	}
	defer errors.Close()

	// multi writer dla błędów złożony z writera do pliku i writera na stdout.
	multi := io.MultiWriter(errors, os.Stderr)

	initLog(ioutil.Discard, os.Stdout, warnings, multi)

	Trace.Println("I have something standard to say.")
	Info.Println("Important Information.")
	Warning.Println("There is something you need to know about.")
	Error.Println("Something has failed.")
}
