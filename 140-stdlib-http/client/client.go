package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Run server first :) to get data in client!
func main() {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Header, "\n")
	fmt.Println(string(body))
}
