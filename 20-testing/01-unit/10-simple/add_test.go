package add

import (
	"testing"
)

func TestAddAddsTwoInts(t *testing.T) {
	result := Add(10, 20)
	if result != 30 {
		t.Errorf("Hohohoho I want 30 but got %d", result)
	}
}

func TestAddAddsThreeInts(t *testing.T) {
	result := Add(10, 20, 30)
	if result != 60 {
		t.Errorf("Hohohoho I want 60 but got %d", result)
	}
}
