// new(T) allocates zeroed storage for a new item of type T and returns its address
// In Go terminology, it returns a pointer to a newly allocated zero value of type T.

// make(T)
// For slices, maps, and channels, make initializes the internal data structure and
// prepares the value for use
package main

import (
	"fmt"
)

type Person struct {
	Name, Address, Notes string
}

func main() {
	a1 := new(Person)
	fmt.Printf("%+v\n", a1) // pointer

	a2 := &Person{}
	a2.Name = "John Doe"
	fmt.Printf("%+v\n", a2) // pointer

	b := new([]int)
	fmt.Printf("%+v\n", b)

	c := make([]int, 10, 100)
	fmt.Printf("%+v\n", c)
}
