package main

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	if right == left || nums[left] < nums[right] {
		return nums[left]
	}

	el := min(nums[left], nums[right])
	for left+1 < right {
		i := (left + right) / 2
		if el == nums[i] {
			n := findMin(nums[left : i+1])
			if n < el {
				el = n
			}
			left = i
			continue
		}
		if el < nums[i] {
			left = i
			continue
		}
		right = i
		el = nums[i]
	}
	return el
}
