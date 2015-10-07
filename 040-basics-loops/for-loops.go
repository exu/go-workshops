package main

import "fmt"

func main() {

	fmt.Println("standard loop")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

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

	fmt.Println("while i < 10")

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	return

	// infinite loop
	fmt.Println("infinite loop")
	i = 0
	for {
		i++
		fmt.Println(i)
	}

	// @todo while

}
