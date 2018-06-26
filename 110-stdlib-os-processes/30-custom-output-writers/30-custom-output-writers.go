package main

import (
	"log"
	"os"
	"os/exec"
)

// Writer implements io.Writer to find more call `go doc io.Writer` in your shell
type Writer struct {
	count int
}

func (writer *Writer) Write(p []byte) (n int, err error) {
	writer.count++
	log.Println(string(p), writer.count)
	return len(p), nil
}

func main() {
	proc := exec.Command("ping", "google.pl")
	proc.Stdout = &Writer{}
	proc.Stderr = &Writer{}
	proc.Start()

	err := proc.Wait()
	ps := proc.ProcessState

	if err != nil {
		log.Println("ERROR", err.Error())
		os.Exit(1)
	} else {
		log.Println("SUCCESS", ps.Success())
	}

}
