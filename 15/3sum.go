package leetcode15

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}

	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)

	for left := 0; left < len(nums)-2; left++ {

		// sum != 0 if lowest number is positive
		if nums[left] > 0 {
			break
		}

		// If left value was previously tested then skip
		if left > 0 && nums[left-1] == nums[left] {
			continue
		}

		mid, right := left+1, len(nums)-1

		// Cycle until mid and right meet
		for mid < right {

			sum := nums[left] + nums[mid] + nums[right]

			// Step either mid or right toward eachother depending on which one is
			// contributed more (+/-) to the non-zero result
			if sum < 0 {
				mid++
			} else if sum > 0 {
				right--
			} else {
				result = append(result, []int{nums[left], nums[mid], nums[right]})

				// Increment past any duplicate values
				for mid < right {
					if nums[mid] == nums[mid+1] {
						mid++
					} else if nums[right] == nums[right-1] {
						right--
					} else {
						break
					}
				}

				// These numbers are solved so iterate away from both
				mid++
				right--
			}
		}
	}
	return result
}
