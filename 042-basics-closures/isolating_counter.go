package main

import (
	"fmt"
)

func sequencer() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
func main() {

	nextInt := sequencer()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := sequencer()
	fmt.Println(newInts())
	fmt.Println(newInts())

}
