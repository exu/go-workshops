package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
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
	// Jak usuniemy poniższego if'a to pozwoli nam to na użycie
	// i usuniemy splita poprzecinku możemy akumulować wartości
	// w stylu:
	//	-deltaT 10s -deltaT 15s
	if len(*i) > 0 {
		return errors.New("Już ustawiłeś flagę interval!")
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

// dla specjalnych typów używamy unkcji Var
var intervalFlag interval

func init() {
	flag.Var(&intervalFlag, "deltaT", "lista okresów czasu po przecinku")
}

func main() {
	flag.Parse()
	fmt.Println(intervalFlag)
}
