// w "Concurent Programming" pojawia się problem dostępu
// do zmiennych, co jeżeli 2 gorutynki uderzą w tym samym czasie
// lub zapiszą wartość? jaka zostanie zapisana.
// Prosty przykład który zobrazuje nam inkrementowanie wartośći
package main

import (
	"fmt"
	// "runtime"
	"time"
)

type intCounter int64

func (c *intCounter) Add(x int64) {
	*c += intCounter(x)
}

func (c *intCounter) Value() (x int64) {
	return int64(*c)
}

func main() {
	counter := intCounter(0)

	for i := 0; i < 100; i++ {
		go func(no int) {
			for i := 0; i < 10000; i++ {
				counter.Add(1)
			}
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println(counter.Value())

}

// Data races - to jest typowy przykład
// aby wykryć automatycznie go udostepnia
// przelacznik -race

// go run -race 0-counter-problem.go

// https://golang.org/doc/articles/race_detector.html
