package astiptr

// Bool transforms a bool into a *bool
func Bool(i bool) *bool {
	return &i
}

// Float transforms a float64 into a *float64
func Float(i float64) *float64 {
	return &i
}

// Int transforms an int into an *int
func Int(i int) *int {
	return &i
}

// Int64 transforms an int64 into an *int64
func Int64(i int64) *int64 {
	return &i
}

// Str transforms a string into a *string
func Str(i string) *string {
	return &i
}
