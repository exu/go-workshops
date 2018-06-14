package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
		syscall.SIGQUIT)

	exit_chan := make(chan int)
	go func() {
		for {
			s := <-signal_chan
			switch s {
			// kill -SIGHUP XXXX
			case syscall.SIGHUP:
				fmt.Println("hungup")
				exit_chan <- 255
			case syscall.SIGKILL:
				fmt.Println("kill'em all")
				// it looks like there is no possibility to catch SIGKILL signal
			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				fmt.Println("Ctrl + C")

			// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				fmt.Println("force stop")
				exit_chan <- 0

			// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				fmt.Println("stop and core dump")
				exit_chan <- 0

			default:
				fmt.Println("Unknown signal.")
				exit_chan <- 1
			}
		}
	}()

	fmt.Printf("Current Process PID: %d\n", os.Getpid())

	code := <-exit_chan
	os.Exit(code)
}
