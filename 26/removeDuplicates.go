package leetcode26

// removeDuplicates takes a sorted array and modifies it in-place so that there
// are no consecutive duplicate values. It returns the length from the start of
// the array that contains no duplicates
func removeDuplicates(nums []int) int {

	// Ensure atleast one value exists
	if len(nums) == 0 {
		return 0
	}

	// Injest the first value
	prev, length := nums[0], 1

	// Overwrite duplicates in nums slice
	for _, n := range nums {
		if n != prev {
			nums[length], prev = n, n
			length++
		}
	}

	return length
}
