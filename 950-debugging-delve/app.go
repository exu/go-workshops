package main

import (
	"fmt"
)

var count int

func loopa() {
	for i := 0; i < 10; i++ {
		fmt.Println("no: ", i)
	}
}

func loopb(c int) {
	for i := 0; i < c; i++ {
		fmt.Println("no from b: ", i)
	}
}

func inner1() {
	inner2()
}

func inner2() {
	inner3()
}

func inner3() {
	loopa()
}

func main() {
	count = 10
	inner1()
	loopa()
	loopa()
	loopb(10)
	loopb(20)
	inner2()
	count = 20
}
