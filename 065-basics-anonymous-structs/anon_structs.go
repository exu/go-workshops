package main

import (
	"fmt"
)

func main() {
	// faster than map[string]interface{}
	// don't need to dynamically alocate memory
	a := struct {
		ID   int
		Name string
	}{1, "Placek"}

	fmt.Println(a)
}
