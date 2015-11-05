package main

import (
	"log"
	"os"
	"os/exec"
)

// Writer implementuje io.Writer w dokumentacji
// proc.Stdout jest taki typ
// `godoc io.Writer`
type Writer struct {
	count int
}

func (writer Writer) Write(p []byte) (n int, err error) {
	writer.count++
	log.Println(string(p), writer.count)
	return len(p), nil
}

func main() {
	proc := exec.Command("/bin/sh", "-c", "ls -la && sleep 2")
	proc.Stdout = Writer{}
	proc.Stderr = Writer{}
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
