// In the previous example we saw how to manage simple
// counter state using atomic operations. For more complex
// state we can use a _[mutex](http://en.wikipedia.org/wiki/Mutual_exclusion)_
// to safely access data across multiple goroutines.

package main

import (
	"fmt"
	"sync"
	"time"
)

type mutexCounter struct {
	mu sync.Mutex
	x  int64
}

func (c *mutexCounter) Add(x int64) {
	c.mu.Lock()
	c.x += x
	c.mu.Unlock()
}

func (c *mutexCounter) Value() (x int64) {
	c.mu.Lock()
	x = c.x
	c.mu.Unlock()
	return
}

func main() {
	counter := mutexCounter{}

	for i := 0; i < 100; i++ {
		go func(no int) {
			for i := 0; i < 10000; i++ {
				counter.Add(1)
			}
		}(i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println(counter.Value())

}
