// slice to "nakładka" na tablicę, trzyma do niej referencję
// jak przypiszesz slice do slice'a to będą wskazywać na tą
// samą tablicę.
// Dogłębnie przeanalizowane w:
// - https://blog.golang.org/slices
// - https://github.com/golang/go/wiki/SliceTricks
// - https://blog.golang.org/go-slices-usage-and-internals
// - http://research.swtch.com/godata
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
