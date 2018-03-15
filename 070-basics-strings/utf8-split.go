package main

import (
	"fmt"
)

func main() {
	// splitting string can kick your asses
	const str = `łódź ⌘ ódź`
	sub := str[:3]
	fmt.Println(sub)

	// use rune FTW!
	runes := []rune(str)
	sub = string(runes[:3])
	fmt.Println(sub)
}
