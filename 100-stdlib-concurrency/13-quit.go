package main

func main() {
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				// Do other stuff
			}
		}
	}()

	// Do stuff

	// Quit goroutine
	quit <- true
}
