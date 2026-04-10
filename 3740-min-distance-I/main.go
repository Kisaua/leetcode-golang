package main

func minimumDistance(nums []int) int {
	numsMap := map[int][]int{}

	for i, num := range nums {
		numsMap[num] = append(numsMap[num], i)
	}

	minD := len(nums) + 1

	for _, v := range numsMap {
		if len(v) > 2 {
			for k := range len(v) - 2 {
				if v[k+2]-v[k] < minD {
					minD = v[k+2] - v[k]
				}
			}
		}
	}
	if minD == len(nums)+1 {
		return -1
	}

	return minD * 2
}
