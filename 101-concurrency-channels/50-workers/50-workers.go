package main

import (
	"fmt"
	"time"
)

// worker otrzyma zadania na kanale jobs i zwróci rezultaty na kanale results.
// worker jest CPU intensive więc jego zadanie trwa całą sekunde :)
// jego super tajnym zadaniem jest podnoszenie liczby do potęgi 2
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		time.Sleep(time.Second)
		fmt.Println("worker", id, "processing job", j)
		results <- j * j
	}
}

func main() {
	// tworzymy kanały
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Startujemy workery, na razie każdy z nich zatrzymał się na pierwszej
	// iteracji ponieważ nie ma jeszcze żadnych danych (i kanał jobs nie został zamknięty)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Wysyłamy 9 zadań do kanału jobs
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs) // zamykamy kanał nie chcemy do niego więcej pisać

	// teraz możemy zebrać wyniki
	for a := 1; a <= 9; a++ {
		//        tu blokujemy odbiór z kanału
		//                                    \
		fmt.Println("result: ", a, " = ", <-results)
		// time.Sleep(time.Second)
	}
}
