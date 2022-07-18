package leetcode1304

func sumZero(n int) []int {
	result := make([]int, n)

	for i := 0; n > 1; i++ {
		result[i] = -n / 2
		result[len(result)-1-i] = n / 2
		n -= 2
	}

	return result
}
