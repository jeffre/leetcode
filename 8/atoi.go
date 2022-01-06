package leetcode8

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// use regexp
// 7743 ns/op
func myAtoi2(s string) int {
	var negative bool
	s = strings.TrimLeft(s, " ")
	r := regexp.MustCompile(`^(\-|\+)?[0-9]+`)
	match := r.FindStringSubmatch(s)

	// No numbers
	if match == nil {
		return 0
	}

	if match[1] == "-" {
		negative = true
	}

	i, err := strconv.Atoi(match[0])
	if err != nil && errors.Is(err, strconv.ErrRange) {
		if negative {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}

	if i > math.MaxInt32 {
		return math.MaxInt32
	} else if i < math.MinInt32 {
		return math.MinInt32
	}

	return i
}

// Following the rules
// 66.99 ns/op
func myAtoi(s string) int {
	runes := []rune(s)
	sign, result, index, n := 1, 0, 0, len(runes)

	// Iterate past spaces from the beginning of the input string
	for index < n && runes[index] == ' ' {
		index++
	}

	// Process any positive or negative markers
	if index < n && runes[index] == '+' {
		sign = 1
		index++
	} else if index < n && runes[index] == '-' {
		sign = -1
		index++
	}

	// Start collecting digits
	for index < n && unicode.IsDigit(runes[index]) {

		// Subtracting two runes to give the integer value between them
		digit := int(runes[index] - '0')

		// Prevent int32 overflow
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > math.MaxInt32%10) {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		// Safe to append new digit
		result = result*10 + digit
		index++
	}

	return sign * result
}
