package main

import "fmt"
import "os"

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}

	// point instance.
	fmt.Printf("%v\n", p)

	// if `p` is struct `+` will add fields information
	fmt.Printf("%+v\n", p)

	// %#v will print source code needed for struct initialisation
	fmt.Printf("%#v\n", p)

	// type
	fmt.Printf("%T\n", p)

	// boolean value
	fmt.Printf("%t\n", true)

	// int (base 10)
	fmt.Printf("%d\n", 123)

	// binary
	fmt.Printf("%b\n", 14)

	// character with given code
	fmt.Printf("%c\n", 33)

	// %x hex
	fmt.Printf("%x\n", 456)

	// floating point numbers - %f.
	fmt.Printf("%f\n", 78.9)

	// %e i %E
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// basic string printing - %s.
	fmt.Printf("%s\n", "\"string\"")

	// double-quote - %q.
	fmt.Printf("%q\n", "\"string\"")

	// %x base-16 string, with two output
	// characters per byte of input.
	fmt.Printf("%x\n", "ó hex this !!")

	// pointers
	fmt.Printf("%p\n", &p)

	// numbers with width
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// floats with width and precision
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// align left
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// strings width
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// strings width aligned left
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// Sprintf returns formatted string
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// Fprintf push strings to io.Writer
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
