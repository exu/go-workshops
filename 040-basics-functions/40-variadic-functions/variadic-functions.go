package main

import "fmt"

// we can receive variadic parameters like
// standard slice
func list(messages ...string) {
	fmt.Println("Items:")
	for _, message := range messages {
		fmt.Println("-", message)
	}
	fmt.Println()
}

func main() {

	list()
	list("coca-cola", "whiskey", "ice")
	list("donuts", "bread")

	// If you already have multiple args in a slice,
	// apply them to a variadic function using
	// func(slice...) like this.
	messages := []string{"A", "B", "C", "D"}
	list(messages...)
}
