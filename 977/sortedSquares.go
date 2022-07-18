package leetcode977

/*
	Given an integer array nums sorted in non-decreasing order, return an array
	of the squares of each number sorted in non-decreasing order.
*/

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
