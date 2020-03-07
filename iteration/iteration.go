package iteration

func Repeat(character string, numberRepeated int) string {
	var repeated string
	for i := 0; i < numberRepeated; i++ {
		repeated += character
	}
	return repeated
}
