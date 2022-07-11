package leetcode209

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

/*
  Use a two-pointer approach. Start both at 0. Increment "right" pointer until
  target is reached. Increment "left" pointer while target is not exceeded.
*/
func minSubArrayLen(target int, nums []int) int {
	minLength := len(nums) + 1
	sum, left := 0, 0

	// Expand right-pointer
	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		// Target satisfied
		for sum >= target && left <= right {
			minLength = min(minLength, right-left+1)

			// Contract left pointer
			sum -= nums[left]
			left++
		}
	}
	if minLength == len(nums)+1 {
		return 0
	}
	return minLength
}
