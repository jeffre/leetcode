package leetcode21

import "fmt"

var debug = false
var step = 1

func generateParenthesis(n int) []string {
	result := []string{}

	if n == 0 {
		return result
	}

	addParenthesis(n, n, "", &result)

	return result
}

// addParenthesis takes a number of open and close parenthesis that need to be
// added to a string (s) and recursively calls itself until all iterations of
// open and close parenthesis are fleshed out
func addParenthesis(open, close int, s string, result *[]string) {

	// If all open and close parenthesis have been used then we have a solution
	// to return
	if open == 0 && close == 0 {
		if debug {
			fmt.Printf("%03d return: %q\n", step, s)
			step++
		}
		*result = append(*result, s)
		return
	}

	// If not all open parenthesis have been used up then append one now
	if open > 0 {
		if debug {
			fmt.Printf("%03d addParenthesis: %d, %d, %q \n", step, open-1, close, s+"(")
			step++
		}
		addParenthesis(open-1, close, s+"(", result)
	}

	// If not all close parenthesis have been used up and there are more opened
	// than closed then we can safely add a close.
	if close > 0 && open < close {
		if debug {
			fmt.Printf("%03d addParenthesis: %d, %d, %q \n", step, open, close-1, s+")")
			step++
		}
		addParenthesis(open, close-1, s+")", result)
	}
}
