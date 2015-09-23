package api

import (
	"testing"
)

func TestGetShoulCallCallFunction(t *testing.T) {
	var callCount = 0
	call = func(method int, url string) {
		callCount++
	}

	Get()
	Get()
	Get()
	Get()

	if callCount != 4 {
		t.Errorf("Shiuld be called 4 times")
	}
}
