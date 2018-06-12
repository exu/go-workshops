package main

import (
	"io"
	"os"
	"os/exec"
)

// go command execution is also stream based
// we can easily pipe commands with use of data pipes
func main() {

	c1 := exec.Command("ls", "/etc")
	c2 := exec.Command("wc", "-l")

	reader, writer := io.Pipe()

	c1.Stdout = writer
	c2.Stdin = reader

	c2.Stdout = os.Stdout

	c1.Start()
	c2.Start()
	c1.Wait()
	writer.Close()
	c2.Wait()
}
