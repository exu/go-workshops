package main

import (
	"fmt"
	"time"
)

func main() {

	// kanał tickera
	tick := time.Tick(100 * time.Millisecond)
	// kanał timera
	boom := time.After(time.Second)

	// bomb! pętla cały czas napierdziela select czeka na
	// wartości z dowolnego kanału, jeżeli ich nie otrzyma
	// wykonuje kod z default:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
