package main

import (
	"time"
)

func main() {
	quit := make(chan bool)
	jobs := make(chan int)

	go func() {
		i := 0
		for {
			i++
			jobs <- i
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for {
			select {
			case <-quit:
				return
			case job := <-jobs:
				println(job)
				time.Sleep(time.Millisecond * 100)
			default:
				println(".")
				time.Sleep(time.Millisecond * 50)
			}
		}
	}()

	// Do stuff
	time.Sleep(time.Second)

	// Quit goroutine
	quit <- true
}
