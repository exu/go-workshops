package main

import "fmt"

func main() {
	// create slice dynamically
	slice := []string{"one", "two", "three"}
	fmt.Printf("Dynamic: %+v\n", slice)
	fmt.Println("Len: ", len(slice))
	fmt.Println("Cap: ", cap(slice))

	// create slice by `make`
	allocated := make([]int, 5, 100)
	fmt.Printf("Allocated: %+v\n", allocated)
	fmt.Println("Len: ", len(allocated))
	fmt.Println("Cap: ", cap(allocated))

}
