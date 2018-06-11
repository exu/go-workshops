package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}

func main() {
	c1 := boring("boring!") // Function returning a channel.
	c2 := boring("boring!") // Function returning a channel.

	for i := 0; i < 10; i++ {
		select {
		case val := <-c1:
			fmt.Printf("You say: %q\n", val)
		case val := <-c2:
			fmt.Printf("You say: %q\n", val)
		}
	}
	fmt.Println("You're boring; I'm leaving.")
}
