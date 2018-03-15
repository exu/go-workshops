package main

import "fmt"

func main() {
	// create slice
	slice := []string{"one", "two", "three"}
	fmt.Println(slice)
	fmt.Println("Len: ", len(slice))
	fmt.Println("Cap: ", cap(slice))
}
