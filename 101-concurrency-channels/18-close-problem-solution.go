package main

import (
	"fmt"
	"sync"
	"time"
)

func gen(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var i int
	for {
		time.Sleep(time.Millisecond * 10)
		ch <- i
		i++
		// when no more data (e.g. from db, or event stream)
		if i > 100 {
			break
		}
	}
}

func receiver(ch chan int) {
	for i := range ch {
		fmt.Println("received:", i)
	}
}

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go gen(ch, wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	receiver(ch)
}
