package main

import (
	"sort"
)

func maxBuilding(n int, restrictions [][]int) int {
	restrictions = append(restrictions, []int{1, 0})
	sort.Slice(restrictions, func(i int, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	})

	if restrictions[len(restrictions)-1][0] != n {
		restrictions = append(restrictions, []int{n, restrictions[len(restrictions)-1][1] + n - restrictions[len(restrictions)-1][0]})
	}

	for i := 1; i < len(restrictions); i++ {
		if restrictions[i-1][1]+restrictions[i][0]-restrictions[i-1][0] < restrictions[i][1] {
			restrictions[i][1] = restrictions[i-1][1] + restrictions[i][0] - restrictions[i-1][0]
		}
	}

	for l := len(restrictions) - 1; l > 0; l-- {
		if restrictions[l][1]+restrictions[l][0]-restrictions[l-1][0] < restrictions[l-1][1] {
			restrictions[l-1][1] = restrictions[l][1] + restrictions[l][0] - restrictions[l-1][0]
		}
	}

	result := 0

	for i := 1; i < len(restrictions); i++ {
		result = max(result, (restrictions[i][1]+restrictions[i-1][1]+restrictions[i][0]-restrictions[i-1][0])/2)
	}

	return result
}
