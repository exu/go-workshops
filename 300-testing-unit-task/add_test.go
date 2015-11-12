package add

import (
	"testing"
)

func assertEquals(t *testing.T, expected int, result int) {
	if result != expected {
		t.Errorf("Hohohoho I want %d but got %d", expected, result)
	}
}

func TestAdd_AddsTwoInts(t *testing.T) {
	assertEquals(t, 30, Add(10, 20))
}

func TestAdd_AddsThreeInts(t *testing.T) {
	assertEquals(t, 60, Add(10, 20, 30))
}

func TestAdd_AddsTenInts(t *testing.T) {
	assertEquals(t, 400, Add(10, 20, 30, 40, 10, 20, 30, 40, 100, 100))
}
