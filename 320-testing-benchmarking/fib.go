package slow

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// func Fib(n int) int {
// 	a, b := 1, 1
// 	for i := 0; i < n; i++ {
// 		a, b = b, a+b
// 	}

// 	return a
// }
