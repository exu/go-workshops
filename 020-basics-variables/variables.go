package main

import "fmt"

func main() {
	// variable declaration with type
	var someNumber int
	someNumber = 1
	fmt.Println("Liczba", someNumber)

	// short assignment without type declaration
	superNumber := 1
	fmt.Println("super liczba", superNumber)

	// Integer max value depends on system architecture!
	var bigNumber int64
	bigNumber = 9223372036854775807

	var intIsInt64On64bitMachines int
	intIsInt64On64bitMachines = 9223372036854775807

	fmt.Println("Ints:", bigNumber, intIsInt64On64bitMachines)

	if bigNumber == int64(intIsInt64On64bitMachines) {
		fmt.Println("bigNumer and intIsInt64On64bitMachines are equal")
	}

	// Array (in go we're using "slices" as interface to arrays - there will be dedicated chapter about arrays)
	someArray := []string{"tab", "licz", "ka"}
	fmt.Println(someArray)

	// Map
	someSimpleMap := map[string]string{"t": "abl", "i": "czka"}
	fmt.Println(someSimpleMap)

	// Some more complicated type
	deepMap := map[string]map[string]map[string]int{
		"t": map[string]map[string]int{
			"a": map[string]int{
				"b": 1,
			},
		},
	}
	fmt.Println(deepMap)
	// definition of above looks like crap
	// but we can fix it for sure using types

	// custom types
	type H map[string]interface{}
	// interface{} means that we want to use any type

	superSimpleDeepMap := H{
		"t": H{
			"a": H{
				"b": 1,
			},
		},
	}

	fmt.Println(superSimpleDeepMap)
}
