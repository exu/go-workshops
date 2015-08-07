package main

import "time"
import "fmt"

func main() {
	fmt.Println("Start")

	timer1 := time.NewTimer(time.Second * 2)

	// Odczyt wartości z timer1.C (channel?) blokuje program
	fmt.Println("Timer 1 expired", <-timer1.C)

	// Do samego sleep'a wystarczy time.Sleep()

	// timery mogą być użyteczne bo możba
	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	time.Sleep(time.Second)
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped before expired")
	}
}

// The first timer will expire ~2s after we start the program, but the second should be
// stopped before it has a chance to expire
