package main

func maxRotateFunction(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	base := 0
	start := 0
	for i := range nums {
		start += i * nums[i]
		base += nums[i]
	}

	maxEl := start
	for i := l - 1; i > 0; i-- {
		start = start + base - l*nums[i]
		maxEl = max(maxEl, start)
	}

	return maxEl
}
