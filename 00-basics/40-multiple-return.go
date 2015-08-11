package main

import "fmt"

// zwracanie multi wartości (definicja w nawiasie, typy po przecinku)
func vals() (int, int) {
	return 42, 2
}

func main() {

	// zwracamy 2 wartości, przypisanie jak w pythonie
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//jeżeli któraś z wartości nas nie interesuje możemy ją
	// pominąć za pomocą znaku _
	_, c := vals()
	fmt.Println(c)
}
