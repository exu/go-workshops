package astitime

import (
	"strconv"
	"time"
)

// Timestamp represents a timestamp
type Timestamp struct {
	time.Time
}

// UnmarshalJSON implements the JSONUnmarshaler interface
func (t *Timestamp) UnmarshalJSON(text []byte) error {
	return t.UnmarshalText(text)
}

// UnmarshalText implements the TextUnmarshaler interface
func (t *Timestamp) UnmarshalText(text []byte) (err error) {
	var i int
	if i, err = strconv.Atoi(string(text)); err != nil {
		return
	}
	t.Time = time.Unix(int64(i), 0)
	return
}

// MarshalJSON implements the JSONMarshaler interface
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return t.MarshalText()
}

// MarshalText implements the TextMarshaler interface
func (t Timestamp) MarshalText() (text []byte, err error) {
	text = []byte(strconv.Itoa(int(t.UTC().Unix())))
	return
}
