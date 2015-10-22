// gotchas
package main

import "fmt"

func main() {
	done := false

	go func() {
		done = true
	}()

	for !done {
	}
	fmt.Println("done!")
}
