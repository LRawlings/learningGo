package main

import "testing"

func TestRomanNumerals(t *testing.T) {

	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
		{"4 gets converted to IV (can't repeat more than 3 times", 4, "IV"},
		{"5 gets converted to V", 5, "V"},
		{"6 gets converted to VI", 6, "VI"},
		{"7 gets converted to VII", 7, "VII"},
		{"8 gets converted to VIII", 8, "VIII"},
		{"9 gets converted to IX", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"19 gets converted to XIX", 19, "XIX"},
		{"23 gets converted to XXIII", 23, "XXIII"},
		{"26 gets converted to XXVI", 26, "XXVI"},
		{"35 gets converted to XXXV", 35, "XXXV"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
		{"59 gets converted to LIX", 59, "LIX"},
		{"78 gets converted to LXXVIII", 78, "LXXVIII"},
		{"89 gets converted to LXXXIX", 89, "LXXXIX"},
		{"90 gets converted to XC", 90, "XC"},
		{"96 gets converted to XCVI", 96, "XCVI"},
		{"100 gets converted to C", 100, "C"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}

}
