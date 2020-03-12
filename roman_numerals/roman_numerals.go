package main

func ConvertToRoman(n int) string {
	switch n {
	case 1:
		return "I"

	case 2:
		return "II"
	case 3:
		return "III"
	default:
		return "not found"
	}

}
