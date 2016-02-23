package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(begin time.Time) {
		fmt.Println("Function call took: ", time.Since(begin))
	}(time.Now())

	time.Sleep(time.Second * 2)
}
