package leetcode10

import (
	"fmt"
	"regexp"
)

// 39.53 ns/op
func isMatch(s string, p string) bool {

	// The only match for an empty pattern is an empty string
	if len(p) == 0 {
		return len(s) == 0
	}

	// set a bool if first characters of p and s match
	firstMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')

	// if wildcard repeat being used
	if len(p) >= 2 && p[1] == '*' {

		// Suppose s="ab" and p=".*"
		// isMatch(s, p[2:])                -> isMatch("ab", "")  -> False
		// firstMatch && isMatch(s[1:], p)  -> isMatch("b", ".*") -> True
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	}

	return firstMatch && isMatch(s[1:], p[1:])
}

// Using regexp
// 7049 ns/op
func isMatch3(s string, p string) bool {
	p = fmt.Sprintf("^%s$", p)
	re := regexp.MustCompile(p)
	return re.MatchString(s)
}
