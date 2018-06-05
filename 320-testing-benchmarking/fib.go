package slow

// In mathematics, the Fibonacci numbers are the numbers in the following integer sequence,
// called the Fibonacci sequence, and characterized by the fact that every number after the
// first two is the sum of the two preceding ones

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// func Fib(n int) int {
// 	a, b := 1, 1
// 	for i := 1; i < n; i++ {
// 		a, b = b, a+b
// 	}

// 	return a
// }
