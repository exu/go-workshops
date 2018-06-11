package main

import "fmt"

func main() {

	fmt.Println("Standard loop")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Print as index and value")
	someArray := []int{1, 2, 3, 4, 5, 6, 7}
	for index, value := range someArray {
		fmt.Println(index, value)
	}

	fmt.Println("Index only loop")
	for index := range someArray {
		fmt.Println(index)
	}

	fmt.Println("Value only loop")
	// If you don't need index you'll need to ignore it explicitly
	for _, value := range someArray {
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
}
