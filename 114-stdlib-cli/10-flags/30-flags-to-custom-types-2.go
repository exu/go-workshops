package main

import (
	"flag"
	"fmt"
	"time"
)

type interval []time.Duration

// Metoda formatująca wartość flagi, część interfejsu flag.Value.
func (i *interval) String() string {
	return fmt.Sprint(*i)
}

// (flag.Value interface).
// tutaj dostajemy wartość z shealla.
// U nas będzie to lista (odzielana przecinkami
// więc trzeba ją pociąć
func (i *interval) Set(value string) error {
	// Możemy używać flag wiele razy, aby zbierać ich wartości
	// w stylu:
	//	-deltaT 10s -deltaT 15s
	duration, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	*i = append(*i, duration)

	return nil
}

// dla specjalnych typów używamy unkcji Var
var intervalFlag interval

func init() {
	flag.Var(&intervalFlag, "deltaT", "lista okresów czasu po przecinku")
}

func echoAfter(message string, durations []time.Duration) {
	for _, interval := range durations {
		time.Sleep(interval)
		fmt.Println(message, "After", interval)
	}
}

func main() {
	flag.Parse()
	echoAfter("Hi", intervalFlag)
}
