package main

import (
	"net/http"

	"github.com/exu/go-workshops/09-project-structure/app/boo-api/handler"
	"github.com/exu/go-workshops/09-project-structure/lib/booer"
)

func main() {
	http.HandleFunc("/", handler.Index)
	http.ListenAndServe(":8080", nil)
}
