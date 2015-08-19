package main

import "fmt"

// `zeroval` ma parametr określony jako `int`
// więc parametr będzie przekazywany jako wartość
// zeroval otrzyma zawsze kopię przekazywanej wartości
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` ma określony parameter w formie `*int`
// co oznacza że przekazywany będzie wskaźnik do zmiennej.
// `*iptr` to skok (dereference) wskaźnika z lokalizacji w
// pamięci do określonej wartości pod tym adresem.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// wywołanie zmiennej w formie `&i` daje nam adres pamięci
	// zmiennej `i`, to jest wskaźnik do `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Wskaźniki mozemy również drukować.
	fmt.Println("pointer:", &i)
}
