package main

import (
	"fmt"
)

func main() {
	array := [...]int{1, 2, 3}

	// przekraczamy poza indeks tablicy
	// zostanie stworzony slice w oparciu
	// o nową tablicę
	slice := array[:3]
	fmt.Println(slice, array)

	slice = append(slice, 10)
	fmt.Println(slice, array)

	slice[0] = 999
	fmt.Println(slice, array)
}
