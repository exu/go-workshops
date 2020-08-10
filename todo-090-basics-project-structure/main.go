package main

import (
	"net/http"

	"myapp/app/boo-api/handler"
)

func main() {
	http.HandleFunc("/", handler.Index)
	http.ListenAndServe(":8080", nil)
}
