package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

// super funkcja
func EvenOdds(x int) string {
	if x%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func validateEvenOdd(x int) bool {
	y := EvenOdds(x)
	fmt.Println(x%2 == 0, y == "even", y, x)
	return x%2 == 0 && y == "even" || x%2 != 0 && y == "odd"
}

func TestEvenOdds(t *testing.T) {

	// napier****
	if err := quick.Check(validateEvenOdd, nil); err != nil {
		t.Error(err)
	}
}
