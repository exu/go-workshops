package main

import (
	"fmt"
	"time"
)

var sem = make(chan int, 5)

func main() {
	for i := 0; i < 20; i++ {
		go handle()
	}

	time.Sleep(time.Second * 10)
}

func handle() {
	sem <- 1
	fmt.Println("Step", time.Now())
	time.Sleep(time.Second)
	<-sem
}
