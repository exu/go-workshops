package main

import (
	"fmt"
	"time"
)

var maxOutstanding = 10
var sem = make(chan int, maxOutstanding)

func main() {
	for {
		go handle()
	}
}

func handle() {
	sem <- 1
	fmt.Println("Step", time.Now())
	time.Sleep(time.Second)
	<-sem
}
