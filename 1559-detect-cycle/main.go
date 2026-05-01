package main

type point struct {
	x, y int
}

type move struct {
	current point
	prev    *point
}

func containsCycle(grid [][]byte) bool {
	visited := map[point]struct{}{}

	l := len(grid[0])
	h := len(grid)

	for i := range h {
		for j := range l {
			if _, ok := visited[point{j, i}]; ok {
				continue
			}

			if isCycle(point{j, i}, visited, grid) {
				return true
			}
		}
	}

	return false
}

func isCycle(start point, visited map[point]struct{}, grid [][]byte) bool {
	moves := []move{{current: start, prev: nil}}
	l := len(grid[0])
	h := len(grid)

	if l < 2 || h < 2 {
		return false
	}
	d := []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for i := 0; i < len(moves); i++ {
		m := moves[i]
		if _, ok := visited[m.current]; ok {
			return true
		}
		visited[m.current] = struct{}{}

		for i := range d {
			if !possibleMove(m, d[i], l, h) {
				continue
			}
			if grid[m.current.y][m.current.x] == grid[m.current.y+d[i].y][m.current.x+d[i].x] {
				moves = append(moves, move{
					current: point{
						m.current.x + d[i].x, m.current.y + d[i].y,
					},
					prev: &m.current,
				})
			}
		}

	}

	return false
}

func possibleMove(m move, direction point, l, h int) bool {
	horizontal := m.current.x + direction.x
	vertical := m.current.y + direction.y

	if horizontal >= l || horizontal < 0 || vertical >= h || vertical < 0 {
		return false
	}

	if m.prev == nil {
		return true
	}

	if horizontal == m.prev.x && vertical == m.prev.y {
		return false
	}

	return true
}
