package main

import (
	"fmt"
	"time"
)

var sem = make(chan int, 5)

func main() {
	for i := 0; i < 20; i++ {
		go handle(i)
	}

	time.Sleep(time.Second * 10)
}

func handle(i int) {
	println("starting", i)
	sem <- 1
	fmt.Println("completed", i, time.Now())
	time.Sleep(time.Second)
	<-sem
}
