package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func handler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 100000000; i++ {
		w.Write([]byte(fmt.Sprintf("%d - %d, ", i, i)))
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
