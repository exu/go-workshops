package slow

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// func Fib(n int) int {
// 	for i, j := 0, 1; j < 100; i, j = i+j, i {
// 		fmt.Println(i)
// 	}
// }
