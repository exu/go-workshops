package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	f1 := flag.NewFlagSet("f1", flag.ContinueOnError)
	silent := f1.Bool("silent", false, "")

	f2 := flag.NewFlagSet("f2", flag.ContinueOnError)
	loud := f2.Bool("loud", false, "")

	switch os.Args[1] {
	case "apply":
		if err := f1.Parse(os.Args[2:]); err == nil {
			fmt.Println("apply", *silent)
		}
	case "reset":
		if err := f2.Parse(os.Args[2:]); err == nil {
			fmt.Println("reset", *loud)
		}
	}
}
