package main

import "fmt"

func main() {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	target := [][]int{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(findRotation(mat, target))
}

func findRotation(mat [][]int, target [][]int) bool {
	same, one, two, three := true, true, true, true

	numRow := len(mat)
	numCol := len(mat[0])

	for i := 0; i < numRow; i++ {
		for j := 0; j < numCol; j++ {
			if one && mat[i][j] != target[j][numRow-1-i] {
				one = false
			}
			if two && mat[i][j] != target[numRow-1-i][numCol-1-j] {
				two = false
			}
			if three && mat[i][j] != target[numCol-1-j][i] {
				three = false
			}
			if same && mat[i][j] != target[i][j] {
				same = false
			}
		}
	}
	return same || one || two || three
}
