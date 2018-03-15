package main

import (
	"fmt"

	// if this code will be pushed to github everyone
	// can import in that way:
	"github.com/exu/go-workshops/010-basics-importing/sub"
	// "./sub" will work too but it's not idiomatic in GO
)

// now we can use some data from our `sub` package
func main() {
	fmt.Println("sub.FirstConstant:\t", sub.FirstContant)
	fmt.Println("sub.Hoł() func call:\t", sub.Hoł())
	fmt.Println("sub.SecondConstant:\t", sub.SecondConstant)
}
