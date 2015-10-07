package main

import (
	"github.com/chzyer/readline"
)

func main() {

	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF
			break
		}

		if line == "help" {
			println(`Usage [run|stop]`)
		}

	}

}
