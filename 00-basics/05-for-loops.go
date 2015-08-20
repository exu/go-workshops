package main

import "fmt"

func main() {

	fmt.Println("index, wartość")
	tablica := []int{1, 2, 3, 4, 5, 6, 7}
	for index, value := range tablica {
		fmt.Println(index, value)
	}

	fmt.Println("sam index")
	for index := range tablica {
		fmt.Println(index)
	}

	fmt.Println("sama wartość")
	// jak nie potrzebujesz indeksu MUSISZ go zignorować
	for _, value := range tablica {
		fmt.Println(value)
	}

	return

	// infinite loop
	i := 0
	for {
		i++
		fmt.Println(i)
	}

	// @todo while

}
