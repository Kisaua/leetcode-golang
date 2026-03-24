package main

import (
	"fmt"
	"math"
)

func main() {
	grid := [][]int{{-1, -2, -3}, {-2, -3, -3}, {-3, -3, -2}}

	fmt.Println(maxProductPath(grid))
}

type point struct {
	x, y int
}

type product struct {
	min, max int
}

type queue []point

func (q *queue) enqueue(value point) {
	*q = append(*q, value)
}

func (q *queue) dequeue() (point, bool) {
	if len(*q) == 0 {
		return point{}, false
	}
	value := (*q)[0]
	*q = (*q)[1:]
	return value, true
}

func maxProductPath(grid [][]int) int {
	result := map[point]product{
		{0, 0}: {
			min: grid[0][0],
			max: grid[0][0],
		},
	}

	que := queue{
		point{0, 0},
	}

	for {
		q, ok := que.dequeue()
		if !ok {
			break
		}

		if q.x+1 < len(grid) {
			p := point{q.x + 1, q.y}
			que.enqueue(p)

			cur := result[q]
			var next *product
			if n, ok := result[p]; ok {
				next = &n
			}
			result[p] = checkPoint(grid[p.y][p.x], cur, next)

		}
		if q.y+1 < len(grid[0]) {
			p := point{q.x, q.y + 1}
			que.enqueue(p)
			cur := result[q]
			var next *product
			if n, ok := result[p]; ok {
				next = &n
			}
			if grid[p.y][p.x] != 0 {
				result[p] = checkPoint(grid[p.y][p.x], cur, next)
			}
		}

	}
	res := result[point{len(grid[0]) - 1, len(grid) - 1}]

	if res.max < 0 {
		return -1
	}

	mod := int(math.Pow10(9) + 7)

	return res.max % mod
}

func checkPoint(val int, cur product, next *product) product {
	ma := cur.min * val
	mi := cur.max * val

	if ma < mi {
		ma, mi = mi, ma
	}

	if next == nil {
		return product{
			min: mi,
			max: ma,
		}
	}

	if next.max < ma {
		next.max = ma
	}

	if next.min > mi {
		next.min = mi
	}

	return *next
}
