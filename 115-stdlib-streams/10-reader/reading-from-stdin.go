package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	buf := bytes.Buffer{}
	buf.ReadFrom(os.Stdin)

	fmt.Printf("%+v", buf)
	fmt.Println("\n\nYou write:", buf.String())
}
