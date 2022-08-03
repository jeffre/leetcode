package leetcode1482

import "math"

/*
	Challenge Description: You are given an integer array bloomDay, an integer m
	and an integer k.

	You want to make m bouquets. To make a bouquet, you need to use k adjacent
	flowers from the garden.

	The garden consists of n flowers, the ith flower will bloom in the
	bloomDay[i] and then can be used in exactly one bouquet.

	Return the minimum number of days you need to wait to be able to make m
	bouquets from the garden. If it is impossible to make m bouquets return -1.
*/

// m is bouquet
// k is adjectency requirement
func minDays(bloomDay []int, m int, k int) int {

	// Ensure quantity of flowers is possible
	if len(bloomDay) < m*k {
		return -1
	}

	// Find the first and last days of bloom
	minDate, maxDate := math.MaxInt, 0
	for _, v := range bloomDay {
		if v > maxDate {
			maxDate = v
		}
		if v < minDate {
			minDate = v
		}
	}

	// Binary search
	for minDate <= maxDate {
		mid := minDate + (maxDate-minDate)/2
		if viable(bloomDay, m, k, mid) {
			maxDate = mid - 1
		} else {
			minDate = mid + 1
		}
	}

	return minDate
}

// m is number of bouquets needed
// k is adjectency of blooms
func viable(bloomDay []int, m, k, day int) bool {
	bouquets := 0
	adjectentBlooms := 0

	// Iterate blooms
	for _, v := range bloomDay {

		if v <= day {
			adjectentBlooms++
		} else {
			adjectentBlooms = 0
		}

		if adjectentBlooms == k {
			adjectentBlooms = 0
			bouquets++
		}
	}
	return bouquets >= m
}
