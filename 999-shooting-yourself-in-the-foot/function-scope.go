package main

import "fmt"

var (
	illKillYou int
)

func init() {
	illKillYou := 1
	illKillYou++
}

func main() {
	fmt.Println(illKillYou)

	other := 10

	{
		other := 999
		fmt.Println("inside block", other)
	}

	fmt.Println("outside block", other)

}

// możesz użyć automatycznego sprawdzania
// go tool vet -shadow 10-function-scope.go
