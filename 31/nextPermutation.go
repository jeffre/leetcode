package leetcode31

import (
	"sort"
)

func nextPermutation(nums []int) {

	// Iterate backwards through nums
	for i := len(nums) - 1; i > 0; i-- {

		// Skip until left-neighbor is smaller
		if nums[i] <= nums[i-1] {
			continue
		}

		// Find the index of the first value larger than num[i] (j). Since the
		// previous step established that the right side of i is sorted this
		// is simple:
		j := len(nums) - 1
		for nums[j] <= nums[i-1] {
			j--
			continue
		}

		// Swap i-1 and j values
		nums[i-1], nums[j] = nums[j], nums[i-1]

		// Sort everything to the right of index
		sort.Ints(nums[i:])

		return
	}

	sort.Ints(nums)
}
