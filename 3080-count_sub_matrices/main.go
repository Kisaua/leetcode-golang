package main

import "fmt"

func main() {
	grid := [][]int{{7, 6, 3}, {6, 6, 1}}

	countSubmatrices(grid, 18)
}

func countSubmatrices(grid [][]int, k int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			var left, top, diag int
			if j != 0 {
				left = grid[i][j-1]
			}
			if i != 0 {
				top = grid[i-1][j]
			}
			if i != 0 && j != 0 {
				diag = grid[i-1][j-1]
			}
			if grid[i][j]+left+top-diag <= k {
				count++
			}
			grid[i][j] += left + top - diag
		}
	}
	return count
}

// prinGrid helper function for debug, prints grids
func printGrid(grid [][]int) {
	for i := range grid {
		fmt.Println(grid[i])
	}
	fmt.Println()
}
