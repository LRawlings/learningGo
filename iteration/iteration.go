package iteration

import "strings"

func Repeat(character string, numberRepeated int) string {
	var repeated string
	for i := 0; i < numberRepeated; i++ {
		repeated += character
	}
	return repeated
}

func hasPrefix(phrase, prefix string) bool {
	return strings.HasPrefix(phrase, prefix)
}

func concatenate(elements []string, separator string) string {
	return strings.Join(elements, separator)
}
