package main

import "fmt"

type Developer struct {
	Hands int
}

func (dev *Developer) CutHand() {
	dev.Hands--
}

func (dev Developer) CutWithSpoon() {
	dev.Hands--
}

func main() {
	dev := Developer{2}
	fmt.Printf("%+v\n", dev)
	fmt.Println("Build na Jenkinsie nie poszedł obcinamy rękę")

	dev.CutHand()
	fmt.Printf("%+v\n", dev)

	dev.CutWithSpoon()
	fmt.Printf("%+v\n", dev)
}
