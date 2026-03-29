package main

import (
	"fmt"
	"math"
)

func main() {
	grid := [][]int{
		{-1, -2, -3},
		{-2, -3, -3},
		{-3, -3, -2},
	}

	fmt.Println(maxProductPath(grid))
}

func maxProductPath(grid [][]int) int {
	minList := make([]int, len(grid[0]))
	maxList := make([]int, len(grid[0]))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			l, t := false, false
			var leftmin, leftmax, topmin, topmax int
			if j != 0 {
				l = true
				leftmin = minList[j-1]
				leftmax = maxList[j-1]
			}
			if i != 0 {
				t = true
				topmin = minList[j]
				topmax = maxList[j]
			}

			minList[j], maxList[j] = minMax(l, t, grid[i][j], leftmin, leftmax, topmin, topmax)
		}
	}

	if maxList[len(grid[0])-1] < 0 {
		return -1
	}
	mod := int(math.Pow10(9) + 7)
	return maxList[len(grid[0])-1] % mod
}

func minMax(l, t bool, cur, leftmin, leftmax, topmin, topmax int) (newmin, newmax int) {
	if !l && !t {
		return cur, cur
	}
	newmin = leftmin * cur
	newmax = leftmax * cur
	if !l {
		newmin = topmin * cur
		newmax = topmax * cur
	}
	if newmin > newmax {
		newmin, newmax = newmax, newmin
	}

	if l {
		lmin := leftmin * cur
		lmax := leftmax * cur
		if lmin > lmax {
			lmin, lmax = lmax, lmin
		}
		newmin = lmin
		newmax = lmax
	}
	if t {
		tmin := topmin * cur
		tmax := topmax * cur
		if tmin > tmax {
			tmin, tmax = tmax, tmin
		}
		if newmin > tmin {
			newmin = tmin
		}
		if newmax < tmax {
			newmax = tmax
		}
	}

	return newmin, newmax
}
