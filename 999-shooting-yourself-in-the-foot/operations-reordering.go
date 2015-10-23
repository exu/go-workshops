package main

import (
	"time"
)

var a, b int

func u1() {
	a = 1
	b = 2
}

func u2() {
	a = 3
	b = 4
}

func p() {
	println(a)
	println(b)
}

func main() {
	go u1()
	go u2()
	go p()
	time.Sleep(100 * time.Millisecond)
}
