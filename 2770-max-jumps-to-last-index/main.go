package main

import "fmt"

func maximumJumps(nums []int, target int) int {
	jumps := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		for j := range i {

			diff := nums[i] - nums[j]
			if diff <= target && diff >= -target {
				if jumps[j]+1 > jumps[i] && (jumps[j] != 0 || j == 0) {
					jumps[i] = jumps[j] + 1
				}
			}
		}

		fmt.Println(jumps)
	}

	if jumps[len(nums)-1] == 0 {
		return -1
	}

	return jumps[len(nums)-1]
}
