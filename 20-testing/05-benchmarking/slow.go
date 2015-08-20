package slow

func ImIdiot(j int) bool {
	return j == 9999999
	// for i := 0; i < 10000000; i++ {
	// 	if j == i && i == 9999999 {
	// 		return true
	// 	}
	// }

	return false
}
