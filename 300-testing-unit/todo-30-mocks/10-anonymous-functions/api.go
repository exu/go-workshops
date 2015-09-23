package api

import (
	"fmt"
)

const (
	// GET method
	GET = iota
	// POST method
	POST
)

var url = "http://api.some.pl/v1/users"

var call = func(method int, url string) {
	fmt.Println(method, url)
}

func Get() {
	call(GET, url)
}

func Post() {
	call(POST, url)
}
