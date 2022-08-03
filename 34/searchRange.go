package leetcode34

/*
	Challenge Description: Given an array of integers nums sorted in
	non-decreasing order, find the starting and ending position of a given
	target value.

	If target is not found in the array, return [-1, -1].

	You must write an algorithm with O(log n) runtime complexity.
*/

func searchRange(nums []int, target int) []int {
	left := search(nums, target, true)
	if left == -1 {
		return []int{-1, -1}
	}

	right := search(nums[left:], target, false)
	return []int{left, left + right}
}

// search performs a binary search that, depending on isFirst, returns first or
// last position of target in nums.
func search(nums []int, target int, wantFirst bool) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2

		// Target found
		if nums[mid] == target {
			if wantFirst {

				// Test for first position of target
				if mid == left || nums[mid-1] != target {
					return mid
				} else {
					right = mid - 1
				}

			} else {

				// Test for last position of target
				if mid == right || nums[mid+1] != target {
					return mid
				} else {
					left = mid + 1
				}

			}
		} else if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// Default to -1 when value not found
	return -1
}
