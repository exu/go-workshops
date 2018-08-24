package main

import (
	"fmt"
	"net/http"

	"github.com/exu/go-workshops/docker/go/image"
)

func handler(w http.ResponseWriter, r *http.Request) {
	image := image.Image{"http://localhost:9999"}
	fmt.Fprintf(w, "Hi there, I love %s!", image.URL)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
