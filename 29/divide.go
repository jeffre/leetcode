package leetcode29

import "math"

/*
	Given two integers dividend and divisor, divide two integers without using
	multiplication, division, and mod operator.

	...

	For this problem, if the quotient is strictly greater than 2^31 - 1, then
	return 2^31 - 1, and if the quotient is strictly less than -2^31, then
	return -2^31.
*/

// Linear solution
func divide(dividend int, divisor int) int {

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend == math.MinInt32 && divisor == 1 {
		return math.MinInt32
	}

	quotient := 0
	dividendPositive, divisorPositive := dividend >= 0, divisor >= 0

	// Change to absolute values
	if !dividendPositive {
		dividend = 0 - dividend
	}
	if !divisorPositive {
		divisor = 0 - divisor
	}

	for ; dividend >= divisor; dividend -= divisor {
		quotient++
	}

	if dividendPositive && divisorPositive || !dividendPositive && !divisorPositive {
		return quotient
	}

	return 0 - quotient
}
