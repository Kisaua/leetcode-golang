package main

func minElement(nums []int) int {
	result := nums[0]

	for i := range nums {
		result = min(result, digitSum(nums[i]))
	}

	return result
}

func digitSum(digit int) int {
	s := 0
	for digit > 0 {
		s += digit % 10
		digit = digit / 10
	}

	return s
}
