package leetcode9

// 17.53 ns/op
func isPalindrome(x int) bool {

	// Negative numbers would results in (eg) `12321-`
	if x < 0 {
		return false
	}

	// A number ending in zero (except zero itself) cannot be a palindrome
	if x%10 == 0 && x != 0 {
		return false
	}

	// Take the last half of the digits of x and invert them;
	// Simultaneously, reduce x to its first half of digits
	inverted := 0
	for inverted < x {
		inverted = inverted*10 + x%10
		x /= 10
	}

	// Check if the two halves match
	// Or in the case of an odd number of digits, we can safely knock the last
	// digit off inverted and test again
	return x == inverted || x == inverted/10
}
