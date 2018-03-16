package main

import (
	"fmt"
	"math/rand"
	"time"
)

func query(out chan int) {
	duration := time.Duration(rand.Int31n(60)) * time.Second

	fmt.Println("Starting ", duration, " query ... ")
	time.Sleep(duration)
	out <- rand.Intn(10000)
}

// Implement timeout from channel. Make it with simpliest way possible, you'll learn
// how to do this better with context package later.

func main() {
	out := make(chan int)

	go func(chan int) {
		for {
			query(out)
		}
	}(out)

	for {
		// implement timeout from
		// query after 1 second
		fmt.Println()
		fmt.Println("Got result", <-out)
	}
}
