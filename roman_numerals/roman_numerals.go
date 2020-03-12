package main

import "strings"

func ConvertToRoman(n int) string {
	var result strings.Builder

	for i := n; i > 0; i-- {
		if i == 4 {
			result.WriteString("IV")
			break
		}
		if i == 5 {
			result.WriteString("V")
			break
		}
		result.WriteString("I")
	}

	return result.String()

}
