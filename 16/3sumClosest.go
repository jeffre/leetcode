package leetcode16

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {

	if len(nums) < 3 {
		panic("not enough nums were provided")
	}

	sort.Ints(nums)

	// Using a safe, arbitrary starting point
	closest := nums[0] + nums[1] + nums[2]

	for left := 0; left < len(nums)-2; left++ {

		// If left value was previously tested then skip
		if left > 0 && nums[left-1] == nums[left] {
			continue
		}

		mid, right := left+1, len(nums)-1

		// Cycle until mid and right meet
		for mid < right {

			currSum := nums[left] + nums[mid] + nums[right]

			// Exact match was found
			if currSum == target {
				return currSum
			}

			// Check if currSum is the closest solution thus far
			if distance(target, currSum) < distance(target, closest) {
				closest = currSum
			}

			// Step either mid or right toward eachother depending on which one
			// can bring us closer to target
			if currSum < target {
				mid++
			} else if currSum > target {
				right--
			}
		}
	}
	return closest
}

func distance(a, b int) int {
	return abs(a - b)
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
