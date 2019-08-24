package iteration

// Repeat takes a character and repeats it the number of times given
func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
