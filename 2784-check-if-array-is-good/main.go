package main

import (
	"slices"
)

func isGood(nums []int) bool {
	slices.Sort(nums)

	l := len(nums)
	if l < 2 || nums[l-1] != l-1 || nums[l-2] != l-1 {
		return false
	}

	for i := range l - 2 {
		if i+1 != nums[i] {
			return false
		}
	}

	return true
}
