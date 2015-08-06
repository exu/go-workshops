package main

import "fmt"

func main() {
	// zminee
	var liczba int
	liczba = 1
	fmt.Println(liczba)

	// krócej
	superliczba := 1
	fmt.Println(superliczba)

	// uwazaj na typy!
	var duzaliczba int64
	duzaliczba = 908132908929083889

	var prawieduzaliczba int
	prawieduzaliczba = 908132908929083889

	fmt.Println(duzaliczba, prawieduzaliczba)

	if duzaliczba == int64(prawieduzaliczba) {
		fmt.Println("Równe są!!")
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

	// hmmm słabe :P

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
