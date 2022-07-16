package leetcode438

type charCount map[byte]int

func (c charCount) has(c2 charCount) bool {
	for b, i := range c {
		if c2[b] != i {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	solutions := []int{}
	if len(p) > len(s) {
		return solutions
	}

	pMap := make(charCount, len(p))
	for _, r := range p {
		pMap[byte(r)]++
	}

	windowMap := make(charCount, len(p))
	for right, r := range s {
		left := right - len(p) + 1
		windowMap[byte(r)]++

		// Populate windowMap to correct size
		if left < 0 {
			continue
		}

		// Decrement the falloff from windowMap
		if left > 0 {
			windowMap[s[left-1]]--
		}

		if pMap.has(windowMap) {
			solutions = append(solutions, left)
		}
	}

	return solutions
}
