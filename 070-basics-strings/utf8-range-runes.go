package main

import (
	"fmt"
)

func main() {
	const nihongo = "日a本ł語ó bijacz!"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}
