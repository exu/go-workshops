package main

import (
	"fmt"

	"github.com/exu/go-workshops/a00-task-http-response-reader/checker"
)

func main() {
	urls := []string{
		"http://google.pl",
		"http://somenotexistingurl.dede/code",
		"http://localhost:8080/api",
		"http://myenglishlab.com",
	}

	results := checker.Check(urls)

	fmt.Println(results)
	//  ___________________
	// < gimme JSON here!! >
	//  -------------------
	//         \   ^__^
	//          \  (oo)\_______
	//             (__)\       )\/\
	//                 ||----w |
	//                 ||     ||

}
