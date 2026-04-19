package main

func maxDistance(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	distance := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			if j-i > distance {
				distance = j - i
			}
			j++
			continue
		}
		i++
		j++

	}

	return distance
}
