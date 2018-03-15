package main

import (
	"fmt"
)

//   / func name
//   |      / parameters
//   |      |                  / return type
//   |      |                 |
func getBoo(boosCount int) string {
	return fmt.Sprintf("There are %d boo(s)", boosCount)
}

// There can be only one main() function in package
func main() {
	println(getBoo(3))
}
