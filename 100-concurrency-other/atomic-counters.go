package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

type atomicCounter struct {
	val int64
}

func (c *atomicCounter) Add(x int64) {
	atomic.AddInt64(&c.val, x)
	runtime.Gosched()
}

func (c *atomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}

func main() {
	counter := atomicCounter{}

	for i := 0; i < 100; i++ {
		go func(no int) {
			for i := 0; i < 10000; i++ {
				counter.Add(1)
			}
		}(i)
	}

	time.Sleep(time.Second)

	fmt.Println(counter.Value())

}
