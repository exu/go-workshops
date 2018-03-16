package main

import (
	"net/http"
)

// handler is used to write data into response body
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Helloł Łerld"))
}

// ListenAndServe start listening for client connections on given port
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
