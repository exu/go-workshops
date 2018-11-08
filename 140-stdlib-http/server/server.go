package main

import (
	"net/http"
)

// handler is used to write data into response body
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Helloł Łerld"))
}
func help(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Use / endpoint Luke! "))
}

// ListenAndServe start listening for client connections on given port
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/help", help)
	http.ListenAndServe(":8080", nil)
}
