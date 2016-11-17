package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Starting command")
	// proc := exec.Command("ssh", "injector@inject02.perf.aws-eu-wi.mel.elt.hosts.pearson-intl.net", "cat", "/home/injector/.bashrc")
	proc := exec.Command("whoami")
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr

	if err := proc.Start(); err != nil {
		panic(err)
	}

	proc.Wait()

	fmt.Println("Command started")
	fmt.Print("Running server: ")
	http.ListenAndServe(":9900", nil)
	fmt.Print(" - DONE!\n")
}
