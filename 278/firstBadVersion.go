package leetcode278

/*
	Challenge Description:
*/

func firstBadVersion(n int) int {
	low, high := 0, n
	for low <= high {
		mid := low + (high-low)/2
		if isBadVersion(mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

// The below is used to emulate the leetcode-given API within our _test.go
var bad = 10001

func isBadVersion(n int) bool {
	return n == bad
}
