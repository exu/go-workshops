package main

import (
	"fmt"
)

type MyInt int

func (myInt *MyInt) AddAnother(num MyInt) {
	*myInt += num
}

func main() {
	num1 := MyInt(1)
	num2 := MyInt(2)

	num1.AddAnother(num2) // MyInt Type
	num1.AddAnother(num2)
	num1.AddAnother(num2)
	num1.AddAnother(num2)
	num1.AddAnother(1000) // internal int

	fmt.Printf("%+v\n", num1)
	fmt.Printf("%+v\n", num2)
}
