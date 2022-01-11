package leetcode17

import (
	"unicode"
)

var phoneLetter = map[int][]string{
	2: {"a", "b", "c"},
	3: {"d", "e", "f"},
	4: {"g", "h", "i"},
	5: {"j", "k", "l"},
	6: {"m", "n", "o"},
	7: {"p", "q", "r", "s"},
	8: {"t", "u", "v"},
	9: {"w", "x", "y", "z"},
}

var phoneLetters = []string{
	"",
	"",
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

// Process through 1 result at a time
// 162246 ns/op
func letterCombinations(digits string) []string {

	result := []string{}

	// Check for valid input
	if len(digits) == 0 {
		return result
	}

	combine(digits, []byte{}, &result)

	return result
}

func combine(digits string, currWord []byte, result *[]string) {

	// If currWord is as long as digits then it is ready to be added to results
	if len(currWord) == len(digits) {
		*result = append(*result, string(currWord))
		return
	}

	// Based on the length of the currWord, get next digit and substract from
	// it the unicode value for '0' leaving an integer 'digit'
	digit := digits[len(currWord)] - '0'

	word := phoneLetters[digit]

	for i := 0; i < len(word); i++ {

		// Append the letter and call self again until the word completes
		combine(digits, append(currWord, word[i]), result)

	}
}

// Process through 1 whole digit at a time
// 318703 ns/op
func letterCombinations2(digits string) []string {

	result := []string{}

	// Iterate over each character in digits
	for _, char := range digits {

		// Valid input must be a digit
		if !unicode.IsDigit(char) {
			break
		}

		digit := int(char - '0')

		// Valid input must be between 2 and 9 (inclusive)
		if digit < 2 || digit > 9 {
			break
		}

		// Create new results based on appending phoneLetter to old results
		// Special case if this is the first iteration
		if len(result) == 0 {
			result = phoneLetter[digit]
		} else {
			updatedResults := []string{}
			for _, s := range result {

				// Iterate through each phoneLetter of a given digit
				for _, l := range phoneLetter[digit] {
					updatedResults = append(updatedResults, s+l)
				}
			}
			result = updatedResults
		}
	}

	return result
}
