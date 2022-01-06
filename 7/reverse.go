package leetcode7

import (
	"math"
)

// 55.84 ns/op
func reverse(x int) int {

	result := 0

	for x != 0 {

		// Separate the ones digit
		pop := x % 10
		x /= 10

		// Prevent int32 overflow
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && pop > math.MaxInt32%10) {
			return 0
		}

		// Prevent int32 underflow
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && pop < math.MinInt32%10) {
			return 0
		}

		// It is now safe to append pop to result
		result = result*10 + pop
	}
	return int(result)
}
