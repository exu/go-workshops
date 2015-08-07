package main

import "log"
import "time"

type Result int

type Conn struct {
	Id int
}

func (c *Conn) DoQuery(params string) Result {
	log.Println("Querying start", params, c.Id)
	time.Sleep(2 * time.Second)
	log.Println("Querying end", params, c.Id)
	return Result(c.Id)
}

func Query(conns []*Conn, query string) Result {
	ch := make(chan Result, 1)
	for _, conn := range conns {
		go func(c Conn) {
			select {
			case ch <- c.DoQuery(query):
			default:
			}
		}(*conn)
	}
	return <-ch
}

func main() {
	conns := []*Conn{&Conn{1}, &Conn{2}}
	result := Query(conns, "gimme one moar time")
	log.Println(result)
}
