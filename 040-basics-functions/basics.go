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

// może być tylko jeden main
func main() {
	println(getBoo(3))
}
