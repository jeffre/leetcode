package leetcode11

// 55166 ns/op
func maxArea(height []int) int {
	first, last := 0, len(height)-1
	maxVolume := 0

	// Start at outter edges of the container and work toward center.
	for first < last {

		// Volume is based on the shorter height
		curVolume := (last - first) * min(height[first], height[last])
		maxVolume = max(curVolume, maxVolume)

		// Increment left or right indexes toward eachother based on height
		if height[first] > height[last] {
			last--
		} else {
			first++
		}
	}
	return maxVolume
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
