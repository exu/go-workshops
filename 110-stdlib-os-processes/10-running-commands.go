package main

import (
	"log"
	"os/exec"
)

func main() {

	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("The date is %s\n", out)
}
