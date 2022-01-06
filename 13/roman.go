package leetcode13

import "strings"

type romanNumeral struct {
	Value  int
	Symbol string
}

type romanNumerals []romanNumeral

// Hold roman numerals in decending order for easy iteration
var allRomanNumerals = romanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func romanToInt(s string) int {
	var result int

	// Cycle over the sorted list of romans
	for _, numeral := range allRomanNumerals {

		// Chip away at roman string (potentially) multiple times
		// and increment its value into result
		for strings.HasPrefix(s, numeral.Symbol) {
			result += numeral.Value
			s = strings.TrimPrefix(s, numeral.Symbol)
		}
	}
	return result
}
