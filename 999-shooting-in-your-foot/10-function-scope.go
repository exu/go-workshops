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
	fmt.Printf("%+v", illKillYou)
}
