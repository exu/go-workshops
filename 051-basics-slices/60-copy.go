package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 3, 4}
	slice2 := make([]int, len(slice))
	copy(slice2, slice)

	slice2[2] = 12391387
	fmt.Println(slice2, slice)

	slice3 := append(slice, []int{}...)
	slice3[2] = 123787847
	fmt.Println(slice3, slice)
}
