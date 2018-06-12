package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	buf := bytes.Buffer{}
	buf.ReadFrom(os.Stdin)

	fmt.Printf("\n\n%+v", buf)
	fmt.Println("\n\nYou write:", buf.String())
}
