package main

import "fmt"
import "time"

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)

	for t := range ticker.C {
		fmt.Println(t)
	}
}
