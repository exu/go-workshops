// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"io"
	"log"
	"os"
)

type PrefixWriter struct{}

func (writer *PrefixWriter) Write(p []byte) (n int, err error) {
	os.Stdout.Write([]byte("[BOO] "))
	os.Stdout.Write(p)
	return len(p) + 6, nil
}

func main() {
	f, _ := os.Create("file.txt")
	defer f.Close()

	prefixWriter := &PrefixWriter{}
	writer := io.MultiWriter(f, os.Stdout, os.Stderr, prefixWriter)

	logger := log.New(writer, "[LOG] ", 0)
	logger.Println("Hahahahah")
}
