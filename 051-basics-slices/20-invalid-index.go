package main

import "fmt"

func main() {

	drinks := make([]string, 5)
	drinks[0] = "Whiskey"
	drinks[1] = "Wódka"
	drinks[2] = "Gin"
	drinks[3] = "Piwo"
	drinks[4] = "Wino"

	drinks[5] = "Odmrażacz szyb samolotowych"

	fmt.Println(drinks)
}
