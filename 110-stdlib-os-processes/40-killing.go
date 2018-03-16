package main

import (
	"log"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("sleep", "12")
	cmd.Start()

	// creating done channel to listen on it when program will complete
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	// listening for available channels data
	select {
	case <-time.After(3 * time.Second):
		if err := cmd.Process.Kill(); err != nil {
			log.Fatal("failed to kill: ", err)
		}
		log.Println("process killed as timeout reached")
	case err := <-done:
		if err != nil {
			log.Printf("process done with error = %v", err)
		} else {
			log.Print("process done gracefully without error")
		}
	}
}
