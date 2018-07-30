package main

import (
	"flag"
	"fmt"
	"time"
)

// One of most powerful part of flag package is definig flags custom types

// assume we need to get from user several time intervals
type interval []time.Duration

// String - method which formats flag value, implementing `flag.Value` interface.
func (i *interval) String() string {
	return fmt.Sprint(*i)
}

// (flag.Value interface).
// Here we get value from shell.
// until now previous example was identical
// but we want to simplify it be removing unnecessary
// complicated part

func (i *interval) Set(value string) error {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	*i = append(*i, duration)

	return nil
}

var intervalFlag interval

func init() {
	flag.Var(&intervalFlag, "d", "time interval (can be repeated many times e.g. -d=1s -d=2s)")
	flag.Parse()
}

func echoAfter(message string, durations []time.Duration) {
	for _, interval := range durations {
		time.Sleep(interval)
		fmt.Println(message, "After", interval)
	}
}

// now use our flag and echo a message with given intervals
func main() {
	echoAfter("Hi", intervalFlag)
}
