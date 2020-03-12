package main

import "strings"

func ConvertToRoman(n int) string {
	var result strings.Builder

	if n == 4 {
		return "IV"
	}

	for i := n; i > 0; i-- {
		if i == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}

	return result.String()

}
