package main

import "fmt"

func main() {
	// simple constant definition
	const MaxCapacity = 5

	// enumerated constant
	const (
		A int = iota
		B
		C
		D
		E
		F
	)

	fmt.Printf("Show me values: A=%d ,B=%d, C=%d, D=%d\n", A, B, C, D)

	// enumerated with arithmetics
	const (
		_          = iota // ignore
		KB float64 = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	fmt.Printf("Bytes %0.0f %0.0f %0.0f %0.0f\n", KB, MB, GB, TB)

	const (
		// Max integer value on 64 bit architecture.
		maxInt = 9223372036854775807

		// Much larger value than int64.
		bigger = 9223372036854775808543522345

		// Will NOT compile
		// biggerint int64 = 9223372036854775808543522345
	)
}
