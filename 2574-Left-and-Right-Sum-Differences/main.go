package main

func leftRightDifference(nums []int) []int {
	l := len(nums)
	result := make([]int, l)
	left := make([]int, l)
	right := make([]int, l)
	if l <= 1 {
		return result
	}
	for i := 1; i < l; i++ {
		left[i] = left[i-1] + nums[i-1]
		right[l-i-1] = right[l-i] + nums[l-i]
	}

	for i := range left {
		if left[i]-right[i] < 0 {
			result[i] = right[i] - left[i]
			continue
		}
		result[i] = -right[i] + left[i]
	}

	return result
}
