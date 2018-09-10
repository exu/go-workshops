// we need channels to pass data between goroutines
package main

import "fmt"
import "time"

func doIt(done chan bool) {
	fmt.Print("working...")
	time.Sleep(2 * time.Second)
	fmt.Println("done")
	// sending value through channel
	done <- true
}

func main() {
	// start goroutine
	done := make(chan bool)
	go doIt(done)

	// wait for value form channel (blocking)
	a := <-done
	fmt.Println(a)
}

// if we remove `<-done` doIt will not be executed because main function will
// end faster than doit
