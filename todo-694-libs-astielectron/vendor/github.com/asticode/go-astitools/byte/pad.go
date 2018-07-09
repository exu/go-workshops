package astibyte

// PadLeft adds n rpl to the left of i so that len is length
func PadLeft(i []byte, rpl byte, length int) []byte {
	if len(i) >= length {
		return i
	} else {
		var o = make([]byte, length)
		o = i
		for idx := 0; idx < length-len(i); idx++ {
			o = append([]byte{rpl}, o...)
		}
		return o
	}
}

// PadRight adds n rpl to the right of i so that len is length
func PadRight(i []byte, rpl byte, length int) []byte {
	if len(i) >= length {
		return i
	} else {
		var o = make([]byte, length)
		o = i
		for idx := 0; idx < length-len(i); idx++ {
			o = append(o, rpl)
		}
		return o
	}
}
