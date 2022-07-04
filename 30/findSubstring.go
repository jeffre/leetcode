package leetcode30

func findSubstring(s string, words []string) []int {

	solutions := []int{}

	wordSize := len(words[0])
	windowSize := len(words) * wordSize

	// Slide window over s from left to right
	for a := 0; a <= len(s)-windowSize; a++ {

		// Track found words
		found := make([]bool, len(words))

		// Break window into word-sized groups
		success := true
		for i := a; i < windowSize+a; i += wordSize {
			subWindow := s[i : i+wordSize]
			if !contains(words, found, subWindow) {
				success = false
				break
			}
		}

		if success {
			solutions = append(solutions, a)
		}
	}

	return solutions
}

func contains(haystack []string, found []bool, needle string) bool {
	for i, h := range haystack {

		// Skip found words
		if found[i] {
			continue
		}

		if h == needle {
			// Mark found
			found[i] = true
			return true
		}
	}
	return false
}
