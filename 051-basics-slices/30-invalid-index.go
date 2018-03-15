package main

import "fmt"

func main() {

	drinks := make([]string, 5)
	drinks[0] = "Whiskey"
	drinks[1] = "Vodka"
	drinks[2] = "Gin"
	drinks[3] = "Wine"
	drinks[4] = "Beer"

	// there will be runtime error (panic) index out of ranges
	drinks[5] = "Water"

	fmt.Println(drinks)
}
