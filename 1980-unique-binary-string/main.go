package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	findDifferentBinaryString([]string{"01", "10"})
}

func findDifferentBinaryString(nums []string) string {
	n := len(nums)

	for i := 0; i < n+1; i++ {
		s := fmt.Sprintf("%%0%ds", n)
		binaryString := fmt.Sprintf(s, strconv.FormatInt(int64(i), 2))

		if !slices.Contains(nums, binaryString) {
			return binaryString
		}

	}

	return ""
}
