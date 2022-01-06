package leetcode12

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

func intToRoman(num int) string {
	var result strings.Builder

	// Cycle over the sorted list of romans
	for _, numeral := range allRomanNumerals {

		// Chip away at num potentially multiple times for each roman
		for num >= numeral.Value {
			result.WriteString(numeral.Symbol)
			num -= numeral.Value
		}
	}

	return result.String()
}
