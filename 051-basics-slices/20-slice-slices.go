package main

import "fmt"

func main() {

	sliceMe := "Ala ma kota"
	fmt.Println("[:]  ", sliceMe)

	slice := sliceMe[2:5]
	fmt.Println("[2:5]", slice)

	slice = sliceMe[:5]
	fmt.Println("[:5] ", slice)

	slice = sliceMe[5:]
	fmt.Println("[5:] ", slice)
}
