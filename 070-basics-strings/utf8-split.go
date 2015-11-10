package main

import (
	"fmt"
)

func main() {
	const str = `łódź ⌘ ódź`
	sub := str[:3]
	fmt.Println(sub)

	runes := []rune(str)
	sub = string(runes[:3])
	fmt.Println(sub)
}
