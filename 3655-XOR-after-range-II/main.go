package main

import (
	"math"
)

func xorAfterQueries(nums []int, queries [][]int) int {
	mod := (int(math.Pow10(9) + 7))
	const chunkSize = 1
	if len(nums) < chunkSize {
		for i := range queries {
			for j := queries[i][0]; j <= queries[i][1]; j += queries[i][2] {
				nums[j] = (nums[j] * queries[i][3]) % mod
			}
		}
		result := 0
		for _, num := range nums {
			result ^= num
		}

		return result
	}

	return 0
}
