// _goroutine_ autorzy określają jako lekki thread.
package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// standardowe wywołanie funkcji (blokujące synchroniczne)
	f("direct")

	// Aby wywołać funkcję synchronicznie dajemy słowo kluczowe `go`
	// `go f(s)`. Ta nowa gorutynka zostanie wywołana w nowym thred'zie
	// równolegle z istniejącą
	go f("goroutine")

	// Możemy wywoływać gorutynki jako funkcje anonimowe
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Do tego miejsca nasze 2 gorutynki działają równolegle na
	// 2 osobnych thread'ach, więc wykonanie programu idzie dalej
	// aby zobaczyć rezultat przyblokujemy dalsze wykonywanie
	var input string
	fmt.Scanln(&input)

	fmt.Println("done")
}
