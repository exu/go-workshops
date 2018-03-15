package main

import (
	"fmt"
)

func main() {
	const strangeString = "日a本ł語ó bijacz!"
	for index, runeValue := range strangeString {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}
