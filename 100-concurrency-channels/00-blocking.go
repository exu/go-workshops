// kanałów używamy do synchronizacji gorutynek
package main

import "fmt"
import "time"

func doIt(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// wysyłanie wartości do kanału
	done <- true
}

func main() {
	// startujemy gorutynke, przekazujemy jej kanał
	done := make(chan bool)
	go doIt(done)

	// blokujemy do otrzymania wartości z tego kanału
	a := <-done
	fmt.Println(a)
}

// jeżeli usuniemy <-done program zostanie zakończony a go worker(done) nie zdąży
// się wykonać.
