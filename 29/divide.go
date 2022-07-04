package leetcode29

import (
	"math"
)

/*
	Given two integers dividend and divisor, divide two integers without using
	multiplication, division, and mod operator.

	...

	For this problem, if the quotient is strictly greater than 2^31 - 1, then
	return 2^31 - 1, and if the quotient is strictly less than -2^31, then
	return -2^31.
*/

// O(logN) 56.84 ns/op
func divide(dividend int, divisor int) int {

	// Prevent int32 overflow
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// Resultant sign
	positive := (dividend < 0) == (divisor < 0)

	// Covert both to negative values. Setting both negative is preferrable to
	// both positive because its absolute value is greater (-2147483648 vs
	// 2147483647).
	if dividend > 0 {
		dividend = 0 - dividend
	}
	if divisor > 0 {
		divisor = 0 - divisor
	}

	// Exponentially grow divisor until it nearly exceeds dividend; call this
	// "accum". And record the mutliplicative difference from divisor to
	// "accum"; call this "multiplier"
	accum := divisor
	multiplier := 1
	for accum >= math.MinInt32>>1 && dividend <= accum+accum {
		multiplier += multiplier
		accum += accum
	}

	// Walk accum back down until it is <= divisor
	quotient := 0
	for dividend <= divisor {
		if dividend <= accum {
			dividend -= accum
			quotient += multiplier
		}
		accum >>= 1
		multiplier >>= 1
	}

	if !positive {
		return 0 - quotient
	}
	return quotient
}
