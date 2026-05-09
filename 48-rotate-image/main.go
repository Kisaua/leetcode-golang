package main

func rotate(matrix [][]int) {
	l := len(matrix)
	for i := range l / 2 {
		for j := i; j < l-1-i; j++ {
			matrix[j][l-1-i], matrix[l-1-i][l-1-j],
				matrix[l-1-j][i], matrix[i][j] = matrix[i][j], matrix[j][l-1-i],
				matrix[l-1-i][l-1-j], matrix[l-1-j][i]
		}
	}
}
