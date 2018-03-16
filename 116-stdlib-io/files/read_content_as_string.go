package main

import (
	"io/ioutil"
)

// sometimes you don't need to stream file
// content and want to read whole into memory
// iouti comes here to help you.
func main() {
	out, err := ioutil.ReadFile("/etc/passwd")
	if err != nil {
		panic(err)
	}

	println(string(out))
}
