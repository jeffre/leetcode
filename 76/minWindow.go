package leetcode76

func minWindow(s string, t string) string {

	// Frequency array has length of 58 to contain the character sets between A-z (specifically:
	// A-Z, 6 unused symbols, and a-z)
	var targetFreq, windowFreq [58]int

	// Window's boundaries
	left, right := 0, -1

	// 'best' variables track the best solution (thus far)
	bestLeft, bestRight, bestLen := -1, -1, len(s)+1

	// The count of characters from target that are within the current window
	matches := 0

	// Populate targetFreq with the count of occurrences for each character
	for i := 0; i < len(t); i++ {
		targetFreq[t[i]-'A']++
	}

	// Iterate through source string
	for left < len(s) {

		// If there is an another character (right+1) in source string
		// AND not enough matches have been found
		if right+1 < len(s) && matches < len(t) {

			// Increment count of occurrences of this source character
			windowFreq[s[right+1]-'A']++

			// If this new character is needed by the target (ie occurs less than or equally in source)
			// than increment 'matches'
			if windowFreq[s[right+1]-'A'] <= targetFreq[s[right+1]-'A'] {
				matches++
			}

			right++

		} else {

			// IF all characters in 't' have been matched
			// AND this window is better than previous bests
			// THEN record a new best
			if matches == len(t) && right-left+1 < bestLen {
				bestLen = right - left + 1
				bestLeft = left
				bestRight = right
			}

			// IF the character on the left of the current window was counted in 'matches'
			// THEN uncount it before 'left' increments forward
			if windowFreq[s[left]-'A'] == targetFreq[s[left]-'A'] {
				matches--
			}

			// Remove this left character from count of occurrences in source string
			windowFreq[s[left]-'A']--

			left++
		}

	}

	if bestLeft == -1 {
		return ""
	} else {
		return s[bestLeft : bestRight+1]
	}
}
