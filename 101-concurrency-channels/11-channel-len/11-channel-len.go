package main

import (
	"fmt"
	"time"
)

func main() {
	// we can get len for buffer in channel
	ch := make(chan int, 100)
	go func() {
		for c := range ch {
			fmt.Println("Readed: ", c)
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 20; i++ {
		ch <- i

		fmt.Printf("Written %d, Len: %d\n", i, len(ch))
	}

	time.Sleep(20 * time.Second)

}
