package main

import (
	"fmt"
)

func main() {
	array := [...]int{1, 2, 3}

	slice := array[:3]
	fmt.Println(slice, array)

	slice = append(slice, 10)
	fmt.Println(slice, array)

	slice[0] = 999
	fmt.Println(slice, array)
}
