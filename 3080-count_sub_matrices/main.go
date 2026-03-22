package main

func main() {
}

func countSubmatrices(grid [][]int, k int) int {
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			left := 0
			top := 0
			if j != 0 {
				left = grid[i][j-1]
			}
			if i != 0 {
				top = grid[i-1][j]
			}
			if grid[i][j]+left+top <= k {
				count++
				grid[i][j] += left + top
			}
		}
	}

	return count
}
