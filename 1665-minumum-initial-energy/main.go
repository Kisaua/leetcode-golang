package main

import (
	"fmt"
	"sort"
)

func minimumEffort(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1]-tasks[i][0] > tasks[j][1]-tasks[j][0]
	})

	fmt.Println(tasks)
	var result, energy int
	for i := range tasks {
		if energy < tasks[i][1] {
			result = result + tasks[i][1] - energy
			energy = tasks[i][1]
		}
		energy -= tasks[i][0]
		fmt.Println(result, energy)
	}
	fmt.Println()
	return result
}
