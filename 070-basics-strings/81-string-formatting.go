package main

import "fmt"
import "os"

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}

	// instancja punktu.
	fmt.Printf("%v\n", p)

	// jeżeli p jest strukturą + doda nam informację o jej polach
	fmt.Printf("%+v\n", p)

	// %#v wyświetli kod źródłowy potrzebny do stworzenia instancji
	fmt.Printf("%#v\n", p)

	// typ wartości
	fmt.Printf("%T\n", p)

	// formatowanie wartości bool
	fmt.Printf("%t\n", true)

	// base 10 dla intów
	fmt.Printf("%d\n", 123)

	// binarna reprezentacja
	fmt.Printf("%b\n", 14)

	// znak o danym kodzie
	fmt.Printf("%c\n", 33)

	// %x hex
	fmt.Printf("%x\n", 456)

	// Zmiennoprzecinkowe
	// %f.
	fmt.Printf("%f\n", 78.9)

	// %e i %E dla formatu naukowego
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// Dla podstawowego drukowania stringów - %s.
	fmt.Printf("%s\n", "\"string\"")

	// double-quote jak w kodzie - %q.
	fmt.Printf("%q\n", "\"string\"")

	// %x base-16 string, with two output
	// characters per byte of input.
	fmt.Printf("%x\n", "ó hex this !!")

	// wskaźniki
	fmt.Printf("%p\n", &p)

	// Szerokość i precyzja liczb
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// Szerokość i precyzja liczb zmiennoprzecinkowych
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// wyrówanie do lewej
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// i dla stringów
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// do lewej
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// Sprintf zwraca zamiast pluć na stdout
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// Fprintf pozwala pluć do dowolnego io.Writer'a
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
