package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

// super funkcja
func EvenOdds(x int64) string {
	if x%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

// nasz walidator
func validateEvenOdd(x int64) bool {
	y := EvenOdds(x)

	fmt.Println("f(", x, ")", y)

	if x%2 == 0 && y == "even" {
		return true
	} else if x%2 != 0 && y == "odd" {
		return true
	} else {
		return false
	}
}

// napier****
func TestEvenOdds(t *testing.T) {
	if err := quick.Check(validateEvenOdd, nil); err != nil {
		t.Error(err)
	}
}
