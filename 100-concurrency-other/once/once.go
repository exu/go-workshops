package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("Starting %d\n", i)
			once.Do(
				func() {
					fmt.Println("Only once", i)
				})
			done <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Println("Result:", <-done)
	}
}
