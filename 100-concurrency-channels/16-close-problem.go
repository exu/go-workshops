package main

import (
	"fmt"
	"time"
)

func gen(ch chan int) {
	var i int
	for {
		time.Sleep(time.Millisecond * 10)
		ch <- i
		i++
		if i > 100 {
			break
		}
	}

	close(ch)
}

func receiver(ch chan int) {
	for i := range ch {
		fmt.Println("received:", i)
	}
}

func main() {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go gen(ch)
	}

	receiver(ch)
}
