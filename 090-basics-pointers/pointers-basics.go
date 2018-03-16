package main

import "fmt"

// `zeroByVal` has parameter with `int` type
// so it'll passed by value, `func zeroByVal` will receive copy
// of this value
func zeroByVal(ival int) {
	ival = 0
}

// `zeroByPointer` has `*int` type
// what means that it'll be passed by reference.
// `*iptr` is pointer dereference to given value in memory.
func zeroByPointer(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroByVal(i)
	fmt.Println("zeroByVal:", i)

	// `&i` gives us memory address of variable `i` (pointer to i)
	zeroByPointer(&i)
	fmt.Println("zeroByPointer:", i)

	// we can print pointers too.
	fmt.Println("pointer:", &i)
}
