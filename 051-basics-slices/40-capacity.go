package main

import "fmt"

func main() {
	// slice is data structure which have array pointer
	// size and capacity
	// capacity relates to array on which slice was made
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	subSlice := slice[:5]
	fmt.Println(slice, subSlice)
	fmt.Println("Capacity of 'slice':", cap(slice), cap(subSlice))

	// subSlice and slice have same pointer to underlying array data structure!!
	subSlice[2] = 10000
	fmt.Println("Slice vs subSlice (note 2nd index is the same)", slice, subSlice)

	// We  can create new slice with same length as previous one
	// and copy variables from main slice
	copied := make([]int, len(subSlice))
	copy(copied, slice)
	copied[3] = 999999
	fmt.Println("Slice copy", slice, copied)

	// Strings are immutable (but works similar to slices)
	str := "Jacek Placek"
	jacek := str[:5]
	jacek = "Wacek"

	fmt.Println(jacek, ",", str)
}
