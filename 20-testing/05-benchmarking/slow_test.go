package slow

import (
	"testing"
)

func TestImIdiot_WithValidData(t *testing.T) {
	result := ImIdiot(9999999)
	if result == false {
		t.Fatal("Should be true!")
	}
}

func TestImIdiot_WithInvalidData(t *testing.T) {
	result := ImIdiot(123)
	if result == true {
		t.Fatal("Should be true!")
	}
}

func BenchmarkImIdiot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ImIdiot(9999999)
	}
}
