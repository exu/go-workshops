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

	// teraz możemy użyć naszego writera np jako writer w logach
	logger := log.New(writer, "\n\n[LOG]", 0)
	logger.Println("Hahahahah")
}
