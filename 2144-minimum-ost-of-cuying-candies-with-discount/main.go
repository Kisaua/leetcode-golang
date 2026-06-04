package main

import "slices"

func minimumCost(cost []int) int {
	slices.Sort(cost)
	total := 0
	for i := range cost {
		if (i+1)%3 != 0 {
			total += cost[len(cost)-1-i]
		}
	}
	return total
}
