package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
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
// we want coma separated list of intervals here
func (i *interval) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("you've set flag interval already")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}

		*i = append(*i, duration)
	}
	return nil
}

var intervalFlag interval

func init() {
	flag.Var(&intervalFlag, "d", "coma separated intervals list (e.g.: 2m,12s,2h,1m")
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
