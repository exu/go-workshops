// variadic functions - mogą być wołane z dowolną liczbą argumentów.
package main

import "fmt"

func list(messages ...string) {
	fmt.Println("Items:")
	for _, message := range messages {
		fmt.Println("-", message)
	}
	fmt.Println()
}

func main() {

	list()
	list("coca-cola", "whiskey", "ice")
	list("pączki", "drożdżówki")

	// możemy przekazać slice poprzez dodanie "..."
	// podczas przekazywanie wartości do funkcji
	messages := []string{"A", "B", "C", "D"}
	list(messages...)
}
