// Elemnty mapy nie są sortowane podczas
// iterowania za pomocą range
package main

import (
	"fmt"
)

func main() {

	warehouse := map[string]int{
		"lizak":   100,
		"guma":    1000,
		"arbuz":   12,
		"cytryna": 22,
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Iteration", i)
		fmt.Println("---------------")
		for k, v := range warehouse {
			fmt.Println(k, v)
		}
		fmt.Println("---------------")
	}
}
