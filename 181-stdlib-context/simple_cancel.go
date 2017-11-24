package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// cancel after new line appears in shell
	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	a(ctx, 5*time.Second, "hello")
}

func a(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
