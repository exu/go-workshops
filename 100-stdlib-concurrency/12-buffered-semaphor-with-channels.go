package main

import (
	"time"
)

var sem = make(chan int, 5)

func main() {
	for i := 0; i < 20; i++ {
		go handle()
	}
}

func handle() {
	sem <- 1
	println("Step", time.Now())
	time.Sleep(time.Second)
	<-sem
}
