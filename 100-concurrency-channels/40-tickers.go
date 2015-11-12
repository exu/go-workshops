package main

import "time"
import "fmt"

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker.C {
			fmt.Println("Ene", t)
		}
	}()

	go func() {
		for t := range ticker.C {
			fmt.Println("Due", t)
		}
	}()

	go func() {
		for t := range ticker.C {
			fmt.Println("Fake", t)
		}
	}()

	// do tej pory mamy 3 gorutynki
	// kazda z nich chce zczytac z kanału tickera wartość
	// poniewaz jest to kanal ktory na razie nie ma wartosci
	// wiec wszstkie zostaja zablokowane

	// pierwsza goroutine ktora pobierze wartosc moze ja wyswietlic
	// reszta jest nadal zblokowana. kolejna dostanie wartosc
	// gdy ticker wygeneruje ja po 500 ms itd...

	time.Sleep(time.Second * 4)

	// rownolegle po 4 sekundach zatrzymujemy program
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

// When we run this program the ticker should tick 3 times before we stop it.
