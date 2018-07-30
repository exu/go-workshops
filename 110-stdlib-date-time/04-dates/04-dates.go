package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	// Creating new instance, always with time zone!
	then := time.Date(
		2018, 06, 17, 20, 34, 58, 651387237, time.UTC)

	p("Whole date:", then)

	// let's get some time parts
	p("year", then.Year())
	p("month", then.Month())
	p("day", then.Day())
	p("hour", then.Hour())
	p("minute", then.Minute())
	p("second", then.Second())
	p("ns", then.Nanosecond())
	p("location", then.Location())

	// also weekday is available
	p("weekday", then.Weekday())

	// We can compare two values each other
	p("is then before now?", then.Before(now))
	p("is then after now?", then.After(now))
	p("is then equal now?", then.Equal(now))

	// we can also subtract one date from another
	// We'll get here very handly type - `Duration`
	diff := now.Sub(then)
	p("difference between then and now =", diff)

	// we can also get parts of interval
	p("- hours", diff.Hours())
	p("- minutes", diff.Minutes())
	p("- seconds", diff.Seconds())
	p("- nanoseconds", diff.Nanoseconds())

	// we can add `Duration` to `Time` objects
	p("adding duration", now.Add(2*diff))
	// or subtract it
	p("subtracting duration", now.Add(-diff))

	// adding using constants in time pkg
	p("adding 2 hours", now.Add(2*time.Hour))
}
