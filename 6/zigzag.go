package leetcode6

import "strings"

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	// Each row gets its own string.Builder
	rows := make([]strings.Builder, numRows)
	curRow := 0
	downward := false

	// Iterate over the string (as runes)
	for _, rune := range s {
		rows[curRow].WriteRune(rune)

		// Oscillate up/downward direction
		if curRow == numRows-1 || curRow == 0 {
			downward = !downward
		}
		if downward {
			curRow++
		} else {
			curRow--
		}
	}

	// Concatenate the rows together
	var result string
	for i := range rows {
		result += rows[i].String()
	}

	return result
}
