package leetcode31

import (
	"sort"
)

func nextPermutation(nums []int) {
	length := len(nums)

	// Early return edge cases
	if length <= 1 {
		return
	}

	// Iterate backwards through nums to find the index (idx) that needs to be
	// increased
	idx := length - 2
	for ; idx > -1; idx-- {
		if nums[idx] < nums[idx+1] {
			break
		}
	}

	// Edge cases
	switch idx {
	case -1:
		// No more permutations are possible; Restart
		sort.IntSlice(nums).Sort()
		return

	case length - 2:
		// Last two nums can be swapped
		nums[idx], nums[idx+1] = nums[idx+1], nums[idx]
		return
	}

	// At this point we know nums[idx+2] exists

	// Of the values to the right of idx, find the index of value that is
	// larger than num[idx] but smaller than the rest
	minIndex := idx + 1
	minNum := nums[minIndex]
	for i := idx + 2; i < length; i++ {
		if nums[i] < minNum && nums[i] > nums[idx] {
			minNum = nums[i]
			minIndex = i
		}
	}

	// Swap idx and minIndex values
	nums[idx], nums[minIndex] = nums[minIndex], nums[idx]

	// Sort everything to the right of index in ascending order
	sort.IntSlice(nums[idx+1:]).Sort()
}
