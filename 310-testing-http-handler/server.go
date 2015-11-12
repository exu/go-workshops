package main

import (
	"net/http"

	"github.com/exu/go-workshops/310-testing-http-handler/handler"
)

func main() {
	http.HandleFunc("/", handler.Simple)
	http.ListenAndServe(":8080", nil)
}
