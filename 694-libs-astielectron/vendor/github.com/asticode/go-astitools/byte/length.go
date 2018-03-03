package astibyte

// ToLength forces the length of a []byte
func ToLength(i []byte, rpl byte, length int) []byte {
	if len(i) == length {
		return i
	} else if len(i) > length {
		return i[:length]
	} else {
		var o = make([]byte, length)
		o = i
		for idx := 0; idx < length-len(i); idx++ {
			o = append(o, rpl)
		}
		return o
	}
}
