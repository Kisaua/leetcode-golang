package main

func maxDistance(colors []int) int {
	distanse := 0
	left := 0
	right := len(colors) - 1

	for distanse < len(colors)-left-1 {
		if colors[left] != colors[right] && distanse < right-left {
			distanse = right - left
			left = left + 1
			right = len(colors) - 1
			continue
		}
		right -= 1
		if left == right {
			left += 1
			right = len(colors) - 1
		}
	}

	return distanse
}

