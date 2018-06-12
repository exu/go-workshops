package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type counter int32

func (c *counter) increment() int32 {
	return atomic.AddInt32((*int32)(c), 1)
}

func (c *counter) get() int32 {
	return atomic.LoadInt32((*int32)(c))
}

func main() {
	var c counter
	quit := make(chan bool)
	jobs := make(chan int32)

	go func() {
		for {
			select {
			case <-quit:
				close(jobs)
				return
			default:
				c.increment()
				jobs <- c.get()
			}
		}
	}()

	go func() {
		for job := range jobs {
			fmt.Printf("process %d\n", job)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// Do stuff
	time.Sleep(time.Second)

	// Quit goroutine
	fmt.Printf("quit on %d\n ", c.get())
	quit <- true
}
