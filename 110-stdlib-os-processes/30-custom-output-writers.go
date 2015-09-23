package main

import (
	"log"
	"os"
	"os/exec"
)

// Writer implementuje io.Writer w dokumentacji
// proc.Stdout jest taki typ
// `godoc io.Writer`
type Writer struct{}

func (writer Writer) Write(p []byte) (n int, err error) {
	log.Println(string(p))
	return len(p), nil
}

func main() {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("The date is %s\n", out)

	proc := exec.Command("/bin/sh", "-c", "ls -la && sleep 2")
	proc.Stdout = Writer{}
	proc.Stderr = Writer{}
	proc.Start()

	err = proc.Wait()
	ps := proc.ProcessState

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	} else {
		log.Println(ps.Success())
	}

}
