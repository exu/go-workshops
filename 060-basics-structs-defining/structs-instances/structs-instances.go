package main

import (
	"fmt"
)

type Thing struct {
	itemsCount int
}

func (thing *Thing) SetCount(count int) {
	thing.itemsCount = count
}

func main() {

	a := Thing{}
	a.SetCount(2)

	b := &Thing{}
	b.SetCount(2)

	c := new(Thing)
	c.SetCount(2)

	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
	fmt.Printf("%+v\n", c)
}
