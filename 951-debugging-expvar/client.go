package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	url := "http://localhost:8080/?user="
	for range time.Tick(time.Duration(rand.Int31n(1000)) * time.Millisecond) {
		users := []string{"kaziu", "maciek", "franek", "hawranek"}
		for _, user := range users {
			fmt.Println("Called", url+user)
			http.Get(url + user)
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
		}
	}
}
