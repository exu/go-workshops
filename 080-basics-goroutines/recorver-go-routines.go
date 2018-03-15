package main

import (
	"fmt"
	"math/rand"
	"time"
)

// recovery from panics example in goroutines

func funcWithPanic(name string) {
	panic("Somethings is not allright at all for worker " + name)
}

func start(name string) {
	defer errorHandler()
	fmt.Println("Worker ", name, "working")

	if num := rand.Intn(5); num == 4 {
		funcWithPanic(name)
	}
}

func errorHandler() {
	if err := recover(); err != nil {
		fmt.Println("Handling panic...")
		fmt.Println("Error message:", err)
	}
}

func main() {
	fmt.Println("Starting")
	for {
		go start("one")
		go start("two")
		go start("three")
		time.Sleep(time.Second)
	}
}
