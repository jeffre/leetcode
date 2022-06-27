package leetcode27

// removeElement removes all occurrences of val in nums in-place.
func removeElement(nums []int, val int) int {

	length := 0

	for _, n := range nums {
		if n == val {
			continue
		} else {
			nums[length] = n
			length++
		}
	}

	return length
}
