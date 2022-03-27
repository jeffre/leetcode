package leetcode1614

func maxDepth(s string) int {
	cur, max := 0, 0

	for _, r := range s {
		if r == '(' {
			cur++
		} else if r == ')' {
			cur--
		}
		if cur > max {
			max = cur
		}
	}
	return max
}
