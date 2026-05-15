package main

func findMin(nums []int) int {
	left := 0
	right := len(nums)
	if right == 1 || nums[left] < nums[right-1] {
		return nums[left]
	}
	el := nums[left]
	for left < right {
		i := (left + right + 1) / 2
		if el < nums[i] {
			left = i
			continue
		}
		right = i - 1
		el = nums[i]
	}
	return el
}
