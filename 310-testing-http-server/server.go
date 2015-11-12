package main

import (
	"net/http"

	"github.com/exu/go-workshops/310-testing-http-server/handler"
)

func main() {
	http.HandleFunc("/", handler.Simple)
	http.ListenAndServe(":8080", nil)
}
