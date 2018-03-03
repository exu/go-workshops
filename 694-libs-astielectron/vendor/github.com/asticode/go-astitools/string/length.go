package astistring

// ToLength forces the length of a string
func ToLength(i, rpl string, length int) string {
	if len(i) == length {
		return i
	} else if len(i) > length {
		return i[:length]
	} else {
		for idx := 0; idx <= length-len(i); idx++ {
			i += rpl
		}
		return i
	}
}
