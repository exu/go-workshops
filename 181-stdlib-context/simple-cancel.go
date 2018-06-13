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
	defer cancel() // call cancel on end of function call

	// cancel after new line appears in shell
	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		ctx.Value("aaa")
		cancel()
	}()

	// block on timer or ctx.Done()
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Timed out!")
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
