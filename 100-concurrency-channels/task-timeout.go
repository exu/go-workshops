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
