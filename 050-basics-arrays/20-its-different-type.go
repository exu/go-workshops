package main

import "fmt"

func main() {
	var five [5]int
	four := [4]int{10, 20, 30, 40}

	five = four

	fmt.Println(four)
	fmt.Println(five)
}
