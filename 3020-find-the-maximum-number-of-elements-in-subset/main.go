package main

func maximumLength(nums []int) int {
	numsMap := map[int]int{}

	for i := range nums {
		if _, ok := numsMap[nums[i]]; !ok {
			numsMap[nums[i]] = 1
			continue
		}
		numsMap[nums[i]]++
	}

	return check(numsMap)
}

func check(nums map[int]int) int {
	result := 1
	for k, v := range nums {
		if k == 1 {
			result = max(result, v-1*(1-v%2))
			continue
		}
		n := sol(nums, k)
		result = max(result, n)

	}

	return result
}

func sol(nums map[int]int, n int) int {
	if nums[n] == 1 {
		return 1
	}
	if _, ok := nums[n*n]; !ok {
		return 1
	}

	return 2 + sol(nums, n*n)
}
