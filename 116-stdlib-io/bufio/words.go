package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	const input = "Chrząszcz brzmi w trzcinie w Szczebrzeszynie,\nI Szczebrzeszyn z tfego słynie.\n"
	// const input = "ą\n"
	scanner := bufio.NewScanner(strings.NewReader(input))

	// We're setting "Split" function,
	// bufio package have some other split functions:
	// ScanWords, ScanRunes, ScanLines, ScanBytes
	scanner.Split(bufio.ScanWords)

	// now we can easily count words in our stream
	// it's of course simple example but we can for example
	// pass here some stream from huge file.
	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)

}
