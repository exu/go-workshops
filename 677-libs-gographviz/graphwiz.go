package main

import (
	"fmt"

	"github.com/awalterschulze/gographviz"
)

func main() {
	graphAst, _ := gographviz.ParseString(`digraph G {}`)
	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		panic(err)
	}
	graph.AddNode("G", "a", nil)
	graph.AddNode("G", "b", nil)
	graph.AddEdge("a", "b", true, nil)
	output := graph.String()

	fmt.Println(output)
}
