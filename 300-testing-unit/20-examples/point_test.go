package point

import "fmt"

func ExamplePoint_String() {
	point := Point{2, 3}
	fmt.Println(point.String())
	// Output:
	// 2:3
}

func ExamplePoint_Move_positive() {
	point := Point{2, 3}
	move := Point{2, 2}
	point.Move(move)
	fmt.Println(point.String())
	// Output:
	// 4:5
}

func ExamplePoint_Move_negative() {
	point := Point{2, 3}
	move := Point{-2, -2}
	point.Move(move)
	fmt.Println(point.String())
	// Output:
	// 0:1
}

// Example functions without output comments are compiled but not executed.

// The naming convention to declare examples for the package, a function F, a type T and
// method M on type T are:

// func Example() { ... }
// func ExampleF() { ... }
// func ExampleT() { ... }
// func ExampleT_M() { ... }

// Multiple example functions for a package/type/function/method may be provided by
// appending a distinct suffix to the name. The suffix must start with a lower-case
// letter.

// func Example_suffix() { ... }
// func ExampleF_suffix() { ... }
// func ExampleT_suffix() { ... }
// func ExampleT_M_suffix() { ... }
