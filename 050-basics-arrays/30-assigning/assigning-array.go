package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{6, 7, 8, 9, 0}

	a = b

	a[0] = 666

	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
}
