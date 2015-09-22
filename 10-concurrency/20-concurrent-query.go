package main

import (
	"fmt"
	"log"
	"time"
)

type Result int

// Nasza struktura połączenia
type Conn struct {
	Id int
}

func (c *Conn) DoQuery(params string) Result {
	log.Println("Querying start", params, c.Id)
	time.Sleep(2 * time.Second)
	log.Println("Querying end", params, c.Id)
	return Result(c.Id)
}

// Multi odpytywacz - wykona równolegle zapytania
// do połączeń zdefiniowanych w tablicy połączeń
func Query(conns []*Conn, query string) []Result {
	ch := make(chan Result, 4)
	for _, conn := range conns {
		go func(c Conn) {
			select {
			case ch <- c.DoQuery(query):
			default:
			}
		}(*conn)
	}

	// blokujemy go routynke odpaloną w pętli
	// zczytujemy wszystkie przesłane wartości z kanałów
	// agregujemy wynik jako tablicę reszultatów
	var result []Result
	for range conns {
		result = append(result, <-ch) // czytamy
	}

	return result
}

// kod
func main() {
	// zdefiniujemy 4 różne połączenia
	conns := []*Conn{&Conn{1}, &Conn{2}, &Conn{3}, &Conn{4}}

	// przekażemy je do naszej funkcji odpytującej
	result := Query(conns, "gimme first time")
	fmt.Println(result)
	// Query(conns, "gimme second time")
	// Query(conns, "gimme one moar time")
}
