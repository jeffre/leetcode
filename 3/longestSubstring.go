package leetcode3

func lengthOfLongestSubstring(s string) int {
	// Interpret s as a list of runes
	runes := []rune(s)

	// Maximum length substring without repeating characters
	longest := 0

	// Map of runes' last known position
	mp := map[rune]int{}

	// Starting point of current substring
	i := 0

	for j, r := range runes {

		// If this rune is already in the map then it has previously occurred
		if val, ok := mp[r]; ok {
			// Update starting point to this new position
			i = max(val, i)
		}
		longest = max(longest, j-i+1)
		mp[r] = j + 1
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
