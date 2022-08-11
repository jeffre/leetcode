package leetcode827

/*
	Challenge Description: You are given an n x n binary matrix grid. You are
	allowed to change at most one 0 to be 1.

	Return the size of the largest island in grid after applying this operation.

	An island is a 4-directionally connected group of 1s.
*/

func largestIsland(grid [][]int) int {
	n := len(grid) //grid is square

	index := 2                 // 0 and 1 are used already
	area := make([]int, n*n+2) // Holds the area of each uniquely index island

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			// Walk through grid and replace genericly named land masses (the
			// 1's) with unique indexes instead. Also store the total area for
			// each island in 'area'.
			if grid[r][c] == 1 {
				area[index] = dfs(grid, r, c, index)
				index++
			}
		}
	}

	// Set default ans to the larsest island incase there are no 0s left to
	// consider flipping.
	ans := 0
	for _, v := range area {
		if v > ans {
			ans = v
		}
	}

	// Check every 0's and consider the sum of its neighbors
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 0 {

				// Record which islands this 0 is touching.
				seen := make(map[int]struct{})
				for _, move := range neighbors(n, r, c) {
					if grid[move/n][move%n] > 1 {
						seen[grid[move/n][move%n]] = struct{}{}
					}
				}

				// bonus represents the new island size we would get after
				// incorporating this tile
				bonus := 1
				for i := range seen {
					bonus += area[i]
				}

				// Record largest solution only
				if bonus > ans {
					ans = bonus
				}
			}
		}
	}

	return ans
}

func dfs(grid [][]int, r, c, index int) int {
	n := len(grid)
	area := 1

	grid[r][c] = index

	for _, move := range neighbors(n, r, c) {
		if grid[move/n][move%n] == 1 {
			area += dfs(grid, move/n, move%n, index)
		}
	}

	return area
}

// neighbors takes the grid length n and coordinates r, c and returns the
// location of viable neighbors. loc, is a single integer that can be
// divided by n to get the row, and modulus n to get the column.
func neighbors(n, r, c int) (loc []int) {
	dr := []int{-1, 0, 1, 0}
	dc := []int{0, -1, 0, 1}
	for k := 0; k < 4; k++ {
		nr := r + dr[k]
		nc := c + dc[k]
		if 0 <= nr && nr < n && 0 <= nc && nc < n {
			loc = append(loc, nr*n+nc)
		}
	}
	return loc
}
