package slow

func ImIdiot(j int) bool {
	for i := 0; i < 10000000; i++ {
		if j == i && i == 9999999 {
			return true
		}
	}

	return false
}
