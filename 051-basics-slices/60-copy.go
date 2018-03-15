package main

import (
	"fmt"
)

func main() {
	// if we want to make changes on copy of slice we need to use `copy` function
	slice := []int{1, 2, 3}
	slice2 := make([]int, len(slice))
	copy(slice2, slice)
	slice2[2] = 4444
	fmt.Println(slice2, slice)

	// look out append returns pointer to first parameter
	slice = []int{1, 2, 3}
	slice3 := append(slice, []int{}...)
	slice3[2] = 5555
	fmt.Println(slice3, slice)

	slice = []int{1, 2, 3}
	slice4 := append([]int{}, slice...)
	slice4[2] = 6666
	fmt.Println(slice4, slice)
}
