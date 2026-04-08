package main

import "math"

func xorAfterQueries(nums []int, queries [][]int) int {
	for i := range queries {
		for j := queries[i][0]; j <= queries[i][1]; j += queries[i][2] {
			nums[j] = (nums[j] * queries[i][3]) % (int(math.Pow10(9) + 7))
		}
	}
	result := 0
	for _, num := range nums {
		result ^= num
	}

	return result
}
