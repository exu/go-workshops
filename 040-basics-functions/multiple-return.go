package main

import "fmt"

// returning many values definition in brackets
func vals() (int, int) {
	return 42, 2
}

func main() {

	// we can return 2 values (or more) at a time
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// if we don't want to use any of values we can bypass it with `_` operator
	_, c := vals()
	fmt.Println(c)

	d, _ := vals()

	println(d)
}
