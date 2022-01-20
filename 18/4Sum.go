package leetcode18

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {

	result := [][]int{}

	if len(nums) < 4 {
		return result
	}

	sort.Ints(nums)

	for a := 0; a <= len(nums)-4; a++ {

		// If 'a' was previously tested then skip doing it again
		if a > 0 && nums[a-1] == nums[a] {
			continue
		}

		for b := a + 1; b <= len(nums)-3; b++ {

			// If 'b' was previously tested then skip doing it again
			if b > a+1 && nums[b-1] == nums[b] {
				continue
			}

			c, d := b+1, len(nums)-1

			// Cycle until 'c' and 'd' meet
			for c < d {

				sum := nums[a] + nums[b] + nums[c] + nums[d]

				// Step either mid or right toward eachother depending on which one is
				// contributed more (+/-) to the non-target

				if sum < target {
					c++
				} else if sum > target {
					d--
				} else {
					result = append(result, []int{nums[a], nums[b], nums[c], nums[d]})

					// Increment past any duplicate values
					for c < d {
						if nums[c] == nums[c+1] {
							c++
						} else if nums[d] == nums[d-1] {
							d--
						} else {
							break
						}
					}

					// These numbers are solved so iterate away from both
					c++
					d--
				}
			}
		}
	}
	return result
}
