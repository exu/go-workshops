package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// make new rand source
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Perm(6)
	fmt.Printf("%v\n", i)

	fmt.Println(" This one will be the same each time script will run ")
	for i := 0; i < 100; i++ {
		fmt.Printf("%+v ", rand.Intn(10))
	}
	fmt.Println("")
}
