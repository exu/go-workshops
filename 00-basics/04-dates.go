package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	// Tworzenie nowej instancji
	// Czas jest zawsze związany ze strefą czasową
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

	p(then)

	// Wyciągamy składowe
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// Weekday
	p(then.Weekday())

	// Porównywanie dwóch wartości
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Interwał pomiędzy dwoma czasami
	// (Zwraca ciekawy typ Duration)
	diff := now.Sub(then)
	p(diff)

	// Możemy zwrócić poszczególne wartości interwału
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Dodajemy Duration
	// by a duration.
	p(then.Add(2 * diff))
	p(then.Add(-diff))
}
