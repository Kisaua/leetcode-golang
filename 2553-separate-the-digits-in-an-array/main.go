package main

func separateDigits(nums []int) []int {
	result := make([]int, 0)
	for _, n := range nums {
		sep := make([]int, 0)
		for n > 0 {
			sep = append(sep, n%10)
			n = n / 10
		}
		for i := range sep {
			result = append(result, sep[len(sep)-1-i])
		}
	}

	return result
}
