package leetcode4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	index1, index2, prevValue, median := 0, 0, 0, 0

	length := len(nums1) + len(nums2)

	// Iterate (just over) halfway into sum of the lengths of both slices using
	// the <= operator
	for i := 0; i <= length/2; i++ {

		// Keep old median value incase there is an even number of nums
		prevValue = median

		if index1 == len(nums1) {
			// All nums1 values have been visited. Focus exclusively on nums2

			median = nums2[index2]
			index2++

		} else if index2 == len(nums2) {
			// All nums2 value have been visited. Focus exclusively on nums1

			median = nums1[index1]
			index1++

		} else if nums1[index1] < nums2[index2] {
			// Iterate past the lower num1 value

			median = nums1[index1]
			index1++

		} else {
			// Iterate past the lower num2 value

			median = nums2[index2]
			index2++
		}
	}

	if length%2 == 0 {
		return float64(prevValue+median) / 2
	}

	return float64(median)
}
