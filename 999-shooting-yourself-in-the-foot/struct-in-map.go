// to jest dziwne :/
package main

type data struct {
	Name string
}

func main() {
	m := map[string]data{"x": {"one"}}
	m["x"].Name = "two" //error
}
