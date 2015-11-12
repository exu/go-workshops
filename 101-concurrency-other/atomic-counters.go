// Podstawowy mechanizm synchronizacji goroutyek to kanały, jednak są też inne:
// poniżej użyjemy atomic counters aby zsynchronizować wartość pomiędzy
// kilkudziesięcioma gorutynami.

package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {

	var ops uint64 // 0 jako domyślna wartość

	// Startujemy gorutynki
	for i := 0; i < 100; i++ {
		go func() {
			for {

				// aby atomowo zinkrementować naszą liczbę operacji
				// dajemy adres do zmiennej za pomocą &
				atomic.AddUint64(&ops, 1)

				// Przepuszczamy pozostałe go rutynki
				runtime.Gosched()

				// In order to ensure that this goroutine doesn’t starve the scheduler, we
				// explicitly yield after each operation with runtime.Gosched(). This
				// yielding is handled automatically with e.g. every channel operation and
				// for blocking calls like time.Sleep, but in this case we need to do it
				// manually.

			}
		}()
	}

	// czekamy 2 sekundy na rezultaty
	time.Sleep(2 * time.Second)

	// Go rutynki jeszcze napierdzielają w ten counter więc jeżeli chcemy bezpiecznie
	// wyciągnąć wartość korzystamu z funkcji atomic.LoadUint64 która
	// wyciągnie kopię aktualnej wartości
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
