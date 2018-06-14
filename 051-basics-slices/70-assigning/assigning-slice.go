package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 0}

	a = b

	a[0] = 666

	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
}
