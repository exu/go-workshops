package main

import "fmt"

func main() {
	// slice to struktura która zawiera wskaźnik do tablicy
	// wielkość slice'a oraz pojemność
	// pojemność slice'a odnosi się do tablicy na podstawie której został stworzony
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	subSlice := slice[:5]
	fmt.Println(slice, subSlice)

	fmt.Println("Capacity of slice:", cap(slice), cap(subSlice))

	// subSlice i slice mają ten sam  pointer do leżącej tablicy!!
	subSlice[2] = 10000
	fmt.Println(slice, subSlice)

	iWantToChange := make([]int, len(subSlice))
	iWantToChange[3] = 999999
	fmt.Println(slice, iWantToChange)

	str := "Jacek Placek"
	jacek := str[:5]
	jacek = "Wacek"

	fmt.Println(jacek, str)
}
