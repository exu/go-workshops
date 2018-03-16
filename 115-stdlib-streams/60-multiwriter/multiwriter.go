// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, _ := os.Create("file.txt")
	defer f.Close()

	writer := io.MultiWriter(f, os.Stdout, os.Stderr)
	// writer.Write([]byte("Hoł hoł hoł"))

	// now we can use our writer e.g. for logging some data in log package
	logger := log.New(writer, "\n\n[LOG]", 0)
	logger.Println("Hahahahah")
}
