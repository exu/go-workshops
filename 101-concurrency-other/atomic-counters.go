// Podstawowy mechanizm synchronizacji goroutyek to kanały, jednak są też inne:
// poniżej użyjemy atomic counters aby zsynchronizować wartość pomiędzy
// kilkudziesięcioma gorutynami.

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type atomicCounter struct {
	val int64
}

func (c *atomicCounter) Add(x int64) {
	atomic.AddInt64(&c.val, x)
	runtime.Gosched()
	// In order to ensure that this goroutine doesn’t starve the scheduler, we
	// explicitly yield after each operation with runtime.Gosched(). This
	// yielding is handled automatically with e.g. every channel operation and
	// for blocking calls like time.Sleep, but in this case we need to do it
	// manually.
}

func (c *atomicCounter) Value() int64 {
	// // Go rutynki jeszcze napierdzielają w ten counter więc jeżeli chcemy bezpiecznie
	// // wyciągnąć wartość korzystamu z funkcji atomic.LoadUint64 która
	// // wyciągnie kopię aktualnej wartości
	return atomic.LoadInt64(&c.val)
}

func main() {
	counter := atomicCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(no int) {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				counter.Add(1)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println(counter.Value())

	// // Startujemy gorutynki
	// for i := 0; i < 100; i++ {
	// 	go func() {
	// 		for {

	// 			// aby atomowo zinkrementować naszą liczbę operacji
	// 			// dajemy adres do zmiennej za pomocą &
	// 			atomic.AddUint64(&ops, 1)

	// 			// Przepuszczamy pozostałe go rutynki
	// 			runtime.Gosched()

	// In order to ensure that this goroutine doesn’t starve the scheduler, we
	// explicitly yield after each operation with runtime.Gosched(). This
	// yielding is handled automatically with e.g. every channel operation and
	// for blocking calls like time.Sleep, but in this case we need to do it
	// manually.

	// 		}
	// 	}()
	// }

	// // czekamy 2 sekundy na rezultaty
	// time.Sleep(2 * time.Second)

	// // Go rutynki jeszcze napierdzielają w ten counter więc jeżeli chcemy bezpiecznie
	// // wyciągnąć wartość korzystamu z funkcji atomic.LoadUint64 która
	// // wyciągnie kopię aktualnej wartości
	// opsFinal := atomic.LoadUint64(&ops)
	// fmt.Println("ops:", opsFinal)
}
