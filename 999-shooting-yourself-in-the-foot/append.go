package main

import (
	"fmt"
)

func main() {
	a := []byte("foo")
	b := append(a, []byte("bar")...)
	c := append(a, []byte("baz")...)

	fmt.Println(string(a), string(b), string(c))
}

// So what’s really happening ? Go will reuse the same underlying array in append() if it can
// do so without resizing the underlying array. So all three of these structs are referencing
// the exact same array in memory. The only practical difference is their length value which
// in the case of a is 3, and in the case of b and c is 6.

// Keep in mind, Go will only reuse the same underlying array if the length of the newly
// created slice is less than or equal to the capacity of the initial slice.

// It’s important to understand what is happening under the hood when using some of the
// built-in slice functions. For more info, Robe Pike wrote a very helpful blog post which
// goes into a lot of useful details around slices.
