package leetcode374

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {

	min, max := 1, n

	for {

		// myGuess is the average of min and max
		myGuess := int(float64(min+max) / 2)

		switch guess(myGuess) {
		case -1:
			// myGuess was too high
			max = myGuess - 1
		case 1:
			// myGuess was too low
			min = myGuess + 1
		case 0:
			// myGuess was correct
			return myGuess
		}
	}
}
