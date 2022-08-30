package leetcode79

/*
	Challenge Description: Given an m x n grid of characters board and a string
	word, return true if word exists in the grid.

	The word can be constructed from letters of sequentially adjacent cells,
	where adjacent cells are horizontally or vertically neighboring. The same
	letter cell may not be used more than once.
*/

func exist(board [][]byte, word string) bool {

	for row := range board {
		for col := range board[row] {
			if backtrack(board, row, col, word) {
				return true
			}
		}
	}

	// no match found
	return false
}

func backtrack(board [][]byte, row, col int, suffix string) bool {
	columns := len(board[0])

	// base case for success
	if len(suffix) == 0 {
		return true
	}

	// check current status
	if board[row][col] != suffix[0] {
		return false
	}

	found := false

	// temporary mask
	board[row][col] = '#'

	for _, move := range neighbors(board, row, col) {
		r, c := move/columns, move%columns
		found = backtrack(board, r, c, suffix[1:])
		if found {
			break
		}
	}

	// undo mask
	board[row][col] = suffix[0]

	return found
}

func neighbors(board [][]byte, r, c int) (loc []int) {
	rows, columns := len(board), len(board[0])
	dr := []int{-1, 0, 1, 0}
	dc := []int{0, -1, 0, 1}
	for i := 0; i < 4; i++ {
		nextRow := r + dr[i]
		nextCol := c + dc[i]
		if 0 <= nextRow && nextRow < rows {
			if 0 <= nextCol && nextCol <= columns {
				loc = append(loc, nextRow*columns+nextCol)
			}
		}
	}
	return loc
}
