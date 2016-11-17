package main

import (
	"fmt"
)

// Isolating data

// Lets say you want to create a function that has access to data that persists even after
// the function exits. For example, you want to count how many times the function has been
// called, or you want to create a fibonacci number generator, but you don't want anyone else
// to have access to that data (so they can't accidentally change it). You can use closures
// to achieve this.

func main() {
	gen := makeFibGen()
	for i := 0; i < 10; i++ {
		fmt.Println(gen())
	}
}

func makeFibGen() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}
}
