// _goroutine_ autorzy określają jako lekki thread.
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// standard function call (synchronous one and blocking)
	f("direct")

	// to call function in goroutine we're using `go` keyword
	go f("goroutine")

	// we can call anonymous functions as in goroutines
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second * 2)

	// until now our 2 function `f` calls are running concurrently.
	// first one in main goroutine second one in new goroutine spawned with `go` keyword.
	// we'll block our main goroutine and wait for all to complete.
}
