package main

import (
	"fmt"
)

func main() {
	var x interface{} = 7 // x has dynamic type int and value 7
	fmt.Printf("%T\n", x)
	i := x.(int) // i has type int and value 7
	fmt.Printf("%T\n", i)

	// type I interface {
	// 	m()
	// }
	//var y I
	// s := y.(string)    // illegal: string does not implement I (missing method m)
	// r := y.(io.Reader) // r has type io.Reader and y must implement both I and io.Reader

}
