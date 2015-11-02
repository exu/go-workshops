package main

import "time"
import "fmt"

func main() {
	// tworzymy 2 kanały
	c1 := make(chan string)
	c2 := make(chan string)

	// każdy z kanałów dostanie wartość po określonym czasie
	// wartość przekazujemy w osobnych go-rutynkach

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// odbiór danych z kanału jest blokujący
	// odbierzemy 3 razy z kanałów
	// - c1
	// - c2
	// - time.After zwraca kanał do którego zostanie
	///             wysłana wartość po określonym czasie
	for i := 0; i < 4; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-time.After(time.Second * 6):
			fmt.Println("timeout after 6 seconds")
		}
	}
}
