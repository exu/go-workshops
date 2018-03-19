package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const port = ":8080"

func greet(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello World! %s from %s", time.Now(), hostname)
}

func main() {
	fmt.Println("Starting server on", port)
	http.HandleFunc("/", greet)
	panic(http.ListenAndServe(port, nil))
}
