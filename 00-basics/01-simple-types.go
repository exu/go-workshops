package main

import "fmt"

func main() {
	// stałe
	const DOOR_NUMBER = 5

	// enumerated
	const (
		_          = iota // ignore
		KB float64 = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	fmt.Println("Bytes", KB, MB, GB, TB)

	const (
		A float64 = iota
		B
		C
		D
		E
		F
	)

	fmt.Println("Kojelno odlicz:", A, B, C, D)

	// zmienne
	var liczba int
	liczba = 1
	fmt.Println("Liczba", liczba)

	// krócej definicja i deklaracja
	superliczba := 1
	fmt.Println("super liczba", superliczba)

	// uwazaj na typy!
	var duzaliczba int64
	duzaliczba = 908132908929083889

	var prawieduzaliczba int
	prawieduzaliczba = 908132908929083889

	fmt.Println("Różne typy", duzaliczba, prawieduzaliczba)

	if duzaliczba == int64(prawieduzaliczba) {
		fmt.Println("Różne typy równe są!!")
	}

	//tablice
	tabliczka := []string{"tab", "licz", "ka"}
	fmt.Println(tabliczka)

	//hashmap
	mapka := map[string]string{"t": "abl", "i": "czka"}
	fmt.Println(mapka)

	// zagnieżdżony
	supermap := map[string]map[string]map[string]int{
		"t": map[string]map[string]int{
			"a": map[string]int{
				"b": 1,
			},
		},
	}
	fmt.Println(supermap)
	// hmmm słabe :P ... ale można lepiej :)

	// custom types
	type H map[string]interface{}
	// interface{} oznacza dowolny typ

	supersimplemap := H{
		"t": H{
			"a": H{
				"b": 1,
			},
		},
	}

	fmt.Println(supersimplemap)
}
