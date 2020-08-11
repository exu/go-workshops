package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 10; i++ {
		go run(ctx, i)
	}

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}

func run(ctx context.Context, i int) {
	t := time.Tick(100 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Got some context.Done(%d)\n", i)
			return
		case <-t:
			fmt.Printf("TICK %d\n", i)
		}
	}
}
