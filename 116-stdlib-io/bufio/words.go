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

	// ustawiamy funkcje split bufio ma ich trochę: ScanWords, ScanRunes, ScanLines, ScanBytes
	scanner.Split(bufio.ScanWords)

	// policzymy słowa
	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)

}
