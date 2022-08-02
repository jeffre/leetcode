package leetcode643

func findMaxAverage(nums []int, k int) float64 {
	var window int

	// populate first window
	for i := 0; i < k; i++ {
		window += nums[i]
	}
	max := window

	// test remaining windows
	for i := k; i < len(nums); i++ {
		window += nums[i] - nums[i-k]
		if window > max {
			max = window
		}
	}

	return float64(max) / float64(k)
}
