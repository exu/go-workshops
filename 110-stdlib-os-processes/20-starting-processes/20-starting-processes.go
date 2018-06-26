package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	proc := exec.Command("ping", "google.pl")
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr
	proc.Start()

	err := proc.Wait()
	ps := proc.ProcessState

	if err != nil {
		log.Println("ERRRRRRR", err.Error())
		os.Exit(1)
	} else {
		log.Println(ps.Success())
	}

}
