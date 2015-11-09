package slow

import "testing"

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib01(b *testing.B) { benchmarkFib(1, b) }
func BenchmarkFib02(b *testing.B) { benchmarkFib(2, b) }
func BenchmarkFib03(b *testing.B) { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }
func BenchmarkFib42(b *testing.B) { benchmarkFib(42, b) }

func Test_Fib(t *testing.T) {
	data := map[int]int{
		1: 1,
		2: 1,
		3: 2,
		4: 3,
		5: 5,
		6: 8,
		7: 13,
		8: 21,
	}

	for input, expected := range data {
		result := Fib(input)
		if result != expected {
			t.Fatalf("Want %d but but %d", expected, result)
		}
	}
}
