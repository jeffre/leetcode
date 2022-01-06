package leetcode1

func twoSum(nums []int, target int) []int {
	for i, num1 := range nums {
		for j, num2 := range nums[i+1:] {
			if num1+num2 == target {
				return []int{i, i + j + 1}
			}
		}
	}
	return []int{}
}
