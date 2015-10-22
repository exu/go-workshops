package main

import (
	"fmt"
)

func main() {
	// szybsze niż map[string]interface{}
	// nie musi dynamicznie alokować pamięci
	a := struct {
		ID   int
		Name string
	}{1, "Placek"}

	fmt.Println(a)
}
