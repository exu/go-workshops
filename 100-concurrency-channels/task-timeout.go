package main

import (
	"fmt"
	"time"
)

func query(out chan int) {
	fmt.Println("Start query")
	time.Sleep(time.Second * 10)
	out <- 1
}

func main() {
	// implement timeout from
	// query after 1 second
}
