package leetcode5

// longestPalindrome uses the "Expand around center" approach
func longestPalindrome(s string) string {

	// Interpret s as a slice of runes
	runes := []rune(s)

	if len(runes) == 0 {
		return ""
	}

	// start and end hold the winning indexes
	start, end := 0, 0

	for i := 0; i < len(runes); i++ {

		// Palindrome centered on a rune
		len1 := expandAroundCenter(runes, i, i)

		// Palindrome centered between two runes
		len2 := expandAroundCenter(runes, i, i+1)

		// Get max of len1 and len2
		length := 0
		if len1 > len2 {
			length = len1
		} else {
			length = len2
		}

		// Check if this is the new leader
		if length > end-start+1 {
			start = i - (length-1)/2
			end = i + length/2
		}
	}

	return string(runes[start : end+1])
}

// expandAroundCenter takes a slice of runes and coordinates (left and right).
// It returns the size of largest palindrome centered on the given coordinates.
func expandAroundCenter(runes []rune, left, right int) int {
	L, R := left, right
	for L >= 0 && R < len(runes) && runes[L] == runes[R] {
		L--
		R++
	}
	return R - L - 1
}

// longestPalindrome2 is my first attempt to solve the problem. In the current
// testcases it is comparably performant to the other ("Expand around center")
// approach
func longestPalindrome2(s string) string {

	// Interpret s as a slice of runes
	runes := []rune(s)

	if len(runes) == 0 {
		return ""
	}

	// Longest discovered palindrome
	longest := []rune{runes[0]}

	for start := 0; start <= len(runes)-1; start++ {
		for end := len(runes); end >= start; end-- {
			s := runes[start:end]

			// Break early if the string isn't longer than the current longest
			// rune
			if len(longest) >= len(s) {
				break
			}

			if isPalindrome(s) {
				longest = s
				break
			}
		}
	}
	return string(longest)
}

func isPalindrome(r []rune) bool {
	l := len(r)
	for i := 0; i < l/2; i++ {
		if r[i] != r[l-i-1] {
			return false
		}
	}
	return true
}
