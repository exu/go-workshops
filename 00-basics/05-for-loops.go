package main

import "fmt"

func main() {

	tablica := []int{1, 2, 3, 4, 5, 6, 7}
	for index, value := range tablica {
		fmt.Println(index, value)
	}

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

}
