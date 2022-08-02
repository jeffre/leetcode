package leetcode349

/*
	Given two integer arrays nums1 and nums2, return an array of their
	intersection. Each element in the result must be unique and you may return
	the result in any order.
*/

// 247.8 ns/op
func intersection(nums1 []int, nums2 []int) []int {
	count := map[int]bool{}
	result := []int{}
	for _, n := range nums1 {
		count[n] = true
	}
	for _, n := range nums2 {
		if count[n] {
			result = append(result, n)
			count[n] = false
		}
	}

	return result
}
