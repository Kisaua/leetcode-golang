package main

func hasValidPath(grid [][]int) bool {
	streetsMap := map[int][]int{
		1: {-1, 0, 1, 0},
		2: {0, -1, 0, 1},
		3: {-1, 0, 0, 1},
		4: {0, 0, 1, 1},
		5: {-1, -1, 0, 0},
		6: {0, -1, 1, 0},
	}

	l := len(grid[0]) - 1
	h := len(grid) - 1
	if l == 0 && h == 0 {
		return true
	}

	begins := streetsMap[grid[0][0]]
	starts := make([][]int, 0)
	for k := range begins {
		if begins[k] == 1 {
			st := make([]int, 4)
			st[k] = 1
			starts = append(starts, st)
		}
	}
	for _, st := range starts {
		start := st
		i, j := 0, 0
	loop:
		for {
			i = i + start[0] + start[2]
			j = j + start[1] + start[3]

			// if i == l && j == h {
			// 	return true
			// }
			if i == 0 && j == 0 {
				break
			}

			if i < 0 || i > l || j < 0 || j > h {
				break
			}
			next := streetsMap[grid[j][i]]
			for k := range 4 {
				n := (k + 2) % 4
				if next[k]*start[n] == -1 {
					if i == l && j == h {
						return true
					}

					next[k] = 0
					start = next
					continue loop
				}
			}
			break
		}
	}
	return false
}
