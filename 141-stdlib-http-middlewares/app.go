package main

import (
	"fmt"
	"net/http"
)

func appVersionMiddleware(handler http.HandlerFunc, version string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-App-Version", version)
		handler.ServeHTTP(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there!")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's very important training lesson!")
}

func main() {
	http.HandleFunc("/hi", appVersionMiddleware(handler, "v3"))
	http.HandleFunc("/about", appVersionMiddleware(about, "v4"))
	http.ListenAndServe(":8080", nil)
}
