package leetcode977

import "sort"

/*
	Given an integer array nums sorted in non-decreasing order, return an array
	of the squares of each number sorted in non-decreasing order.
*/

// Trivial solution O(n*log(n))
func sortedSquares2(nums []int) []int {
	for i := range nums {
		nums[i] *= nums[i]
	}
	sort.Ints(nums)
	return nums
}

// Faster solution O(n)
func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	results := make([]int, len(nums))
	for i := len(nums) - 1; left <= right; i-- {
		var square int
		if -nums[left] > nums[right] {
			square = nums[left] * nums[left]
			left++
		} else {
			square = nums[right] * nums[right]
			right--
		}
		results[i] = square
	}
	return results
}
