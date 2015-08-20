package main

import (
	"net/http"

	"github.com/exu/go-workshops/09-project-structure/lib/booer"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	w.Write([]byte(booer.JSON))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
