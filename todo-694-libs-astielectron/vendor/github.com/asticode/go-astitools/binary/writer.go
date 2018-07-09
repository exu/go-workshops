package astibinary

import (
	"fmt"
	"math"
)

// Writer represents a writer
type Writer struct {
	b             []byte
	remainingBits string
}

// New represents a binary writer
func New() *Writer {
	return &Writer{b: []byte{}}
}

// Bytes returns the writer bytes
func (w *Writer) Bytes() []byte {
	return w.b
}

// Write writes binary stuff
func (w *Writer) Write(i interface{}) {
	var s string
	switch i.(type) {
	case string:
		s = i.(string)
	case []byte:
		for _, b := range i.([]byte) {
			s += fmt.Sprintf("%.8b", b)
		}
	case bool:
		if i.(bool) {
			s = "1"
		} else {
			s = "0"
		}
	case uint8:
		s = fmt.Sprintf("%.8b", i)
	case uint16:
		s = fmt.Sprintf("%.16b", i)
	case uint32:
		s = fmt.Sprintf("%.32b", i)
	case uint64:
		s = fmt.Sprintf("%.64b", i)
	default:
		return
	}

	for _, c := range s {
		w.remainingBits = w.remainingBits + string(c)
		if len(w.remainingBits) == 8 {
			var nb byte
			var power float64
			for idx := len(w.remainingBits) - 1; idx >= 0; idx-- {
				if w.remainingBits[idx] == '1' {
					nb = nb + uint8(math.Pow(2, power))
				}
				power++
			}
			w.b = append(w.b, nb)
			w.remainingBits = ""
		}
	}
}

// Reset resets the writer
func (w *Writer) Reset() {
	w.b = []byte{}
	w.remainingBits = ""
}
