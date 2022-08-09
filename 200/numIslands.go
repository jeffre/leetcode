package leetcode200

/*
	Challenge Description: Given an m x n 2D binary grid grid which represents a
	map of '1's (land) and '0's (water), return the number of islands.

	An island is surrounded by water and is formed by connecting adjacent lands
	horizontally or vertically. You may assume all four edges of the grid are
	all surrounded by water.
*/

func numIslands(grid [][]byte) int {
	nr, nc := len(grid), len(grid[0])
	islands := 0

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == 49 {
				islands++
				zero(grid, r, c)
			}
		}
	}

	return islands
}

// zero takes a grid and zeros-outs the parts of it that are contiguous to
// row,col using DFS recursion.
func zero(grid [][]byte, r, c int) {
	nr, nc := len(grid), len(grid[0])

	// Work is done. The last condition is where the magic is at. It prevents
	// the zeroing-out from jumping to another island.
	if r < 0 || c < 0 || r >= nr || c >= nc || grid[r][c] == 48 {
		return
	}

	// Zero-out the given position and all 4 other sides
	grid[r][c] = 48
	zero(grid, r-1, c)
	zero(grid, r+1, c)
	zero(grid, r, c-1)
	zero(grid, r, c+1)
}
