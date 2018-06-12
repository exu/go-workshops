package main

import "fmt"

// slicing slices example
func main() {

	sliceMe := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("[]  ", sliceMe)

	slice := sliceMe[2:5]
	fmt.Println("[2:5]", slice)

	slice = sliceMe[:5]
	fmt.Println("[:5] ", slice)

	slice = sliceMe[5:]
	fmt.Println("[5:] ", slice)

	slice = sliceMe[:]
	fmt.Println("[:] ", slice)
}
