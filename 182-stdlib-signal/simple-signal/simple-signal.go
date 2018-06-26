package main

import (
	"fmt"
	"os"
	"os/signal"
	"time" // or "runtime"
)

func cleanup() {
	fmt.Println("Starting cleanup ... ")
	time.Sleep(time.Second)
	fmt.Println(".... Cleanup completed")
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		fmt.Println("Got signal: ", <-c)
		cleanup()
		os.Exit(1)
	}()

	fmt.Println("I'm some daemon and I'm:")
	for {
		fmt.Println("working hard ...")
		time.Sleep(10 * time.Second) // or runtime.Gosched() or similar per @misterbee
	}

	// try to run go run signal.go
	// C-c
	// try to kill `killall signal`
}
