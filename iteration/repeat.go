package iteration

// Repeat repeats a chars n amount of times and returns the concatenate string.
func Repeat(char string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += char
	}

	return repeated
}
