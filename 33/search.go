package leetcode33

/*
	Challenge Description: There is an integer array nums sorted in ascending
	order (with distinct values).

	Prior to being passed to your function, nums is possibly rotated at an
	unknown pivot index k (1 <= k < nums.length) such that the resulting array
	is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]
	(0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3
	and become [4,5,6,7,0,1,2].

	Given the array nums after the possible rotation and an integer target,
	return the index of target if it is in nums, or -1 if it is not in nums.

	You must write an algorithm with O(log n) runtime complexity.
*/

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {

			// Problem solved
			return mid

		} else if nums[left] <= nums[mid] {

			// Left-half is not rotated. We can ignore all of right-half.
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else {

			// Left-half is rotated. Ergo, right-half is not. Thus we pivot
			// based on right-half values instead.
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}

		}
	}

	// Default case for no solution found
	return -1
}
