package main

import (
	"fmt"
	"time"
)

var maxOutstanding = 5
var sem = make(chan int, maxOutstanding)

func main() {
	for i := 0; i < 20; i++ {
		go handle()
	}

	time.Sleep(time.Minute)
}

func handle() {
	sem <- 1
	fmt.Println("Step", time.Now())
	time.Sleep(time.Second)
	<-sem
}
