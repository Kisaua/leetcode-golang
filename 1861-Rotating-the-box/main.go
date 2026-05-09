package main

func rotateTheBox(boxGrid [][]byte) [][]byte {
	l := len(boxGrid[0])
	h := len(boxGrid)

	result := make([][]byte, l)
	for i := range result {
		result[i] = make([]byte, h)
	}

	for i := range h {
		freeCell := l - 1
		for j := l - 1; j >= 0; j-- {
			s := boxGrid[i][j]
			switch s {
			case '.':
				result[j][h-1-i] = '.'
			case '*':
				freeCell = j - 1
				result[j][h-1-i] = '*'
			case '#':
				result[j][h-1-i] = '.'
				result[freeCell][h-1-i] = '#'
				freeCell -= 1
			}
		}
	}

	return result
}
