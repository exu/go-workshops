package astiflag

import (
	"time"
)

// Time represents a time flag
type Time struct {
	time.Time
}

// Set implements the flag.Value interface
func (f *Time) Set(i string) (err error) {
	f.Time, err = time.Parse(time.RFC3339, i)
	return
}
