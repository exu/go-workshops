package main

import (
	"time"
)

func main() {
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				println(".")
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	// Do stuff
	time.Sleep(time.Second)

	// Quit goroutine
	quit <- true
}
