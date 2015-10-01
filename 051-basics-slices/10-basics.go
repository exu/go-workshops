package main

import "fmt"

func main() {
	// tworzenie slice'ów
	slice := []string{"jeden", "dwa", "trzy"}
	fmt.Println(slice)
	fmt.Println("Len: ", len(slice))
	fmt.Println("Cap: ", cap(slice))
}
