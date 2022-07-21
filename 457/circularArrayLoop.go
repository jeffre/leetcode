package leetcode457

func circularArrayLoop(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	visited := make([]bool, len(nums))

	for i := range nums {
		if visited[i] {
			continue
		}
		visited[i] = true

		// Iterate until a cycle is found
		slow, fast := i, i
		for {
			slow = next(nums, slow)
			visited[slow] = true

			for k := 0; k < 2; k++ {
				fast = next(nums, fast)
				visited[fast] = true
			}

			if slow == fast {
				break
			}
		}

		// Find start of cycle by interating another time from the starting
		// position until it meets where slow/fast first met.
		start := i
		for start != slow {
			start = next(nums, start)
			slow = next(nums, slow)
		}

		// Ensure cycle length is > 1
		if start == next(nums, start) {
			continue
		}

		// Ensure one-direction
		isForward := nums[start] > 0
		n := next(nums, start)
		for n != start {
			if (nums[n] > 0) != isForward {
				// Abort cycle prematurely if direction changed
				break
			}
			n = next(nums, n)
		}
		if n == start {
			return true
		}
	}
	return false
}

func next(nums []int, curPos int) int {
	n := curPos + nums[curPos]
	n %= len(nums)
	if n < 0 {
		n += len(nums)
	}
	return n
}
