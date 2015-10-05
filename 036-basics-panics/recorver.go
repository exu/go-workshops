package main

import (
	"fmt"
)

func start() {
	panic("Somethings is not allright at all\n")
}

func errorHandler() {
	if err := recover(); err != nil {
		fmt.Println("Handling panic...")
		fmt.Println("Error message:", err)
	}
}

func main() {
	fmt.Println("Starting")
	defer errorHandler()
	start()
	fmt.Println("Stopping")
}
