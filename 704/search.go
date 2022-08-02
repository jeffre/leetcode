package leetcode704

/*
	Given an array of integers nums which is sorted in ascending order, and an
	integer target, write a function to search target in nums. If target exists,
	then return its index. Otherwise, return -1.

	You must write an algorithm with O(log n) runtime complexity.

*/

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	// binary search
	for left <= right {
		//mid := (left + right) / 2 // This would overflow if len(nums) is big
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
