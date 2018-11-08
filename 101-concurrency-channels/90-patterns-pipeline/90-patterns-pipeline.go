package main

import (
	"fmt"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			time.Sleep(time.Millisecond * 100)
			out <- n * n
		}
		close(out)
	}()
	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			time.Sleep(time.Millisecond * 100)
			out <- n + n
		}
		close(out)
	}()
	return out
}

func main() {
	c := gen(10, 20, 30, 40, 50)
	out := double(double(c))

	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)

	fmt.Println("Second generator")
	for n := range double(square(square(gen(2, 3, 4, 5, 10, 20, 50, 100)))) {
		fmt.Println(n)
	}
}
