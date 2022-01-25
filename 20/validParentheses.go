package leetcode20

// 272.6 ns/op
// Solved using recursion
func isValid(s string) bool {

	// No parenthesis is valid, but just one is not.
	switch len(s) {
	case 0:
		return true
	case 1:
		return false
	}

	// Find all top level open and close parentheses
	var unclosed, start int
	var opener, closer rune
	for i, r := range s {

		if unclosed == 0 {
			unclosed = 1
			start = i
			opener = r
			switch r {
			case '(':
				closer = ')'
			case '{':
				closer = '}'
			case '[':
				closer = ']'
			}
			continue
		}

		switch r {
		case closer:
			unclosed--
		case opener:
			unclosed++
		}

		if unclosed == 0 {
			// Recursively check the validity of whats inside the parentheses
			if !isValid(s[start+1 : i]) {
				return false
			}
		}
	}
	return unclosed == 0
}

// 347.7 ns/op
// Solved using FILO
func isValid2(s string) bool {

	// Create a first-in, last-out stack
	var stack []rune

	for _, r := range s {

		switch r {

		// If the rune is an opening parentheses, append the appropriate
		// closing parentheses into the FILO stack
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')

		// Otherwise, the rune must be a closing parentheses that needs to be
		// validated
		default:

			// If no parentheses are currently opened then closing is invalid
			if len(stack) == 0 {
				return false
			}

			// Ensure the type of parenthesis being closed matches the one that
			// was opened
			if r != stack[len(stack)-1] {
				return false
			}

			// Passing the above tests, we can safely prune the last item off
			// the stack
			stack = stack[:len(stack)-1]
		}

	}

	// Finally, ensure there are not any (open) parenthesis left in the stack
	return len(stack) == 0
}
