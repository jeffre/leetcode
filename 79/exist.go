package leetcode79

/*
	Challenge Description: Given an m x n grid of characters board and a string
	word, return true if word exists in the grid.

	The word can be constructed from letters of sequentially adjacent cells,
	where adjacent cells are horizontally or vertically neighboring. The same
	letter cell may not be used more than once.
*/

type grid [][]byte

func exist(board [][]byte, word string) bool {

	grid := grid(board)

	for row := range grid {
		for col := range grid[row] {
			if grid.backtrack(row, col, word) {
				return true
			}
		}
	}

	// no match found
	return false
}

func (g grid) backtrack(row, col int, suffix string) bool {
	columns := len(g[0])

	// base case for success
	if len(suffix) == 0 {
		return true
	}

	// check current status
	if g[row][col] != suffix[0] {
		return false
	}

	found := false

	// temporary mask
	g[row][col] = '#'

	for _, move := range g.neighbors(row, col) {
		r, c := move/columns, move%columns
		found = g.backtrack(r, c, suffix[1:])
		if found {
			break
		}
	}

	// undo mask
	g[row][col] = suffix[0]

	return found
}

func (g grid) neighbors(r, c int) (loc []int) {
	rows, columns := len(g), len(g[0])
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
