package main

import (
	"io/ioutil"
)

func main() {
	out, err := ioutil.ReadFile("/etc/passwd")
	if err != nil {
		panic(err)
	}

	println(string(out))
}
