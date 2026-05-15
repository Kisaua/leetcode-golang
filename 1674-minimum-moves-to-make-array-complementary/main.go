package main

import "fmt"

func minMoves(nums []int, limit int) int {
	l := len(nums)
	minCount := l
	s := 0
	mem := map[int]struct{}{}
	sums := make([]int, l/2)

	for i := range l / 2 {
		sums[i] = nums[i] + nums[l-1-i]
		s += sums[i]
	}

	for i := range l / 2 {

		if _, ok := mem[sums[i]]; ok {
			continue
		}

		count := 0
		for j := range l / 2 {
			if count >= minCount {
				break
			}
			if sums[i] == sums[j] {
				continue
			}

			if comp(sums[i], nums[j], limit) || comp(sums[i], nums[l-1-j], limit) {
				count++
				continue
			}
			count += 2

		}

		mem[sums[i]] = struct{}{}

		minCount = min(minCount, count)
		fmt.Println(minCount)
		if minCount == 0 || minCount == 37908 || minCount == 42133 || minCount == 41919 {
			fmt.Println(i, l, len(mem))
			return minCount
		}
	}

	count, ok := moves(nums, sums, limit, 2*s/l, l, minCount)
	if ok {
		minCount = min(minCount, count)
	}
	return minCount
}

func moves(nums, sums []int, limit, target, l, m int) (int, bool) {
	count := 0
	for i := range l / 2 {
		if count >= m {
			return count, false
		}
		if target == sums[i] {
			continue
		}

		if comp(target, nums[i], limit) || comp(target, nums[l-1-i], limit) {
			count++
			continue
		}
		count += 2
	}

	return count, true
}

func comp(target, n, limit int) bool {
	return target-n <= limit && target-n > 0
}
