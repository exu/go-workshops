package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Printf("handler started")
	defer log.Printf("handler ended")

	// run some long term action
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "HellO")
	// or simply listen to context end
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
