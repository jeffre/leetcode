package leetcode340

/*
	Given a string s and an integer k, return the length of the longest
	substring of s that contains at most k distinct characters.
*/

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 || len(s) == 0 {
		return 0
	}

	maxLength := 1
	left, right := 0, 0
	charPos := make(map[byte]int, len(s))

	for right < len(s) {
		charPos[s[right]] = right

		// len(charPos) reprensets the number of unique characters
		if len(charPos) == k+1 {

			// Find lowest index in current window
			lowest := len(s) + 1
			for _, v := range charPos {
				if v < lowest {
					lowest = v
				}
			}
			delete(charPos, s[lowest])
			left = lowest + 1
		}

		maxLength = max(maxLength, right-left+1)
		right++
	}

	return maxLength
}
