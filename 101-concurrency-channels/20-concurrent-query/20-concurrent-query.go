package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Result int

// Nasza struktura połączenia
type Conn struct {
	Id int
}

func (c *Conn) DoQuery(params string) Result {
	log.Println("Querying start", params, c.Id)
	time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
	log.Println("Querying end", params, c.Id)

	return Result(1000 + c.Id*c.Id)
}

// Multi odpytywacz - wykona równolegle zapytania
// do połączeń zdefiniowanych w tablicy połączeń
func Query(conns []Conn, query string) Result {
	ch := make(chan Result)
	for _, conn := range conns {
		go func(c Conn) {
			ch <- c.DoQuery(query)
		}(conn)
	}

	return <-ch
}

// kod
func main() {
	// zdefiniujemy 4 różne połączenia
	conns := []Conn{Conn{1}, Conn{2}, Conn{3}, Conn{4}, Conn{5}}

	// przekażemy je do naszej funkcji odpytującej
	result := Query(conns, "query!")
	fmt.Println(result)
}
