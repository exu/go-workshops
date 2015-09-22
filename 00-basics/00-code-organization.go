package main

import (
	"fmt"

	// jeżeli kod będzie na githubie (lub w innym public repo)
	// każdy może
	// zaimportować w ten sposób
	"github.com/exu/go-workshops/00-basics/sub"

	// TIP: zadziała również
	// import "./sub"
)

func main() {
	fmt.Println("sub.FIRST_CONSTANT:\t", sub.FIRST_CONSTANT)
	fmt.Println("sub.SECOND_CONSTANT:\t", sub.SECOND_CONSTANT)
	fmt.Println("sub.Hoł() func call:\t", sub.Hoł())
}
