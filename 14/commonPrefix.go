package leetcode14

import (
	"strings"
	"unicode/utf8"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	// Iterate strings
	for i := 1; i < len(strs); i++ {

		// Iterate runes of string
		for !strings.HasPrefix(strs[i], prefix) {
			r, size := utf8.DecodeLastRuneInString(prefix)

			// An empty prefix returns (RuneError, 0)
			// An invalid utf-8 character in prefix returns (RuneError, 1)
			if r == utf8.RuneError && (size == 0 || size == 1) {
				size = 0
			}

			// Shorten prefix each iteration
			prefix = prefix[:len(prefix)-size]
		}
	}

	return prefix
}
