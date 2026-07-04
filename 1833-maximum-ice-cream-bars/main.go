package main

import (
	"slices"
)

func maxIceCream(costs []int, coins int) int {
	slices.Sort(costs)

	for i := range costs {
		coins -= costs[i]
		if coins < 0 {
			return i
		}
	}

	return len(costs)
}
