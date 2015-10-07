package main

import (
	"fmt"
)

//                              named return variables
//                            /
func getBoo(boosCount int) (name string, boosCountPlusOne int) {
	boosCountPlusOne = boosCount + 1
	name = fmt.Sprintf("There are %d boo(s)", boosCountPlusOne)
	// return name (could be more)
	return
}

func main() {
	name, count := getBoo(3)
	println(name)
	println(count)
}
