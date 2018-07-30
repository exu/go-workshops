package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 50)

	for i := 1; i <= 50; i++ {
		requests <- i
		// fmt.Println(len(requests), cap(requests))
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 1000)

	// Przez zablokowanie requestów przez kanał tickera
	// limitujemy się do wykonania przebiegu pętli
	// co 200 ms
	for req := range requests {
		fmt.Println("request", req, <-limiter)
	}

}

// spróbuj zmienić program aby działał bez użycia buforu w kanale
