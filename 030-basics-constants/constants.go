package main

import "fmt"

func main() {
	// stałe
	const DOOR_NUMBER = 5

	// enumerated
	const (
		A float64 = iota
		B
		C
		D
		E
		F
	)

	fmt.Println("Kojelno odlicz:", A, B, C, D)

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

	fmt.Println("Bytes", KB, MB, GB, TB)

	const (
		// Max integer value on 64 bit architecture.
		maxInt = 9223372036854775807

		// Much larger value than int64.
		bigger = 9223372036854775808543522345

		// Will NOT compile
		// biggerint int64 = 9223372036854775808543522345
	)
}
