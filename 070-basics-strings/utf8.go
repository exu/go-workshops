package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("Helloł, 世界")

	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c %v\n", r, size)

		b = b[size:]
	}

	// FullRune example
	fmt.Println()
	fmt.Println()
	buf := []byte{228, 184, 150} // 世
	fmt.Println(utf8.FullRune(buf))
	fmt.Println(utf8.FullRune(buf[:2]))

}
