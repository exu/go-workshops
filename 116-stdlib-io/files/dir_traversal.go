package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// path/filepath can recursively Walk in whole
// directory tree and is quite fast, we can
// pass anonymous function to be called for each
// node
func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}
