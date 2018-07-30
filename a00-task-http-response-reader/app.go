package main

import (
	"fmt"

	"./checker"
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
