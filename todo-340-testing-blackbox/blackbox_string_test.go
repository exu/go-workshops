package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

func IsDuck(animal string) bool {
	if animal == "duck" {
		return true
	} else {
		return false
	}
}

func TestIsDuck(t *testing.T) {
	err := quick.Check(
		func(input string) bool { // hmm closure
			result := IsDuck(input)
			fmt.Println("f(", input, ")", result)

			return input != "duck" && !result
		},
		nil,
	)

	if err != nil {
		t.Error(err)
	}
}
