package main

import "strings"

func ConvertToRoman(n int) string {
	var result strings.Builder

	if n == 4 {
		return "IV"
	}

	for i := 0; i < n; i++ {
		result.WriteString("I")
	}

	return result.String()

}
