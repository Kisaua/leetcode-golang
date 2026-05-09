package main

func rotateGrid(grid [][]int, k int) [][]int {
	l := len(grid[0])
	h := len(grid)

	for i := 0; i < l/2 && i < h/2; i++ {
		linelength := 2 * (l + h - 4*i - 2)
		line := make([]int, linelength)
		for j := 0; j < l-2*i; j++ {
			line[j] = grid[i][j+i]
			line[2*(l-2*i)+h-2*i-2-j-1] = grid[h-1-i][j+i]
		}
		for j := 0; j < h-2*i-2; j++ {
			line[l-2*i+j] = grid[j+1+i][l-1-i]
			line[linelength-1-j] = grid[j+1+i][i]
		}

		shifted := append(line[k%linelength:], line[:k%linelength]...)
		for j := 0; j < l-2*i; j++ {
			grid[i][j+i] = shifted[j]
			grid[h-1-i][j+i] = shifted[2*(l-2*i)+h-2*i-2-j-1]
		}
		for j := 0; j < h-2*i-2; j++ {
			grid[j+1+i][l-1-i] = shifted[l-2*i+j]
			grid[j+1+i][i] = shifted[linelength-1-j]
		}
	}
	return grid
}
