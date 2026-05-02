package main

import (
	"slices"
	"strconv"
	"strings"
)

func rotatedDigits(n int) int {
	count := 0

	for i := range n + 1 {
		if isGoodV1(i) {
			count++
		}
	}

	return count
}

func isGood(n int) bool {
	numS := strconv.Itoa(n)

	numSlice := strings.Split(numS, "")

	if slices.Contains(numSlice, "3") || slices.Contains(numSlice, "4") || slices.Contains(numSlice, "7") {
		return false
	}

	rotationMap := map[string]string{
		"0": "0",
		"1": "1",
		"2": "5",
		"5": "2",
		"6": "9",
		"8": "8",
		"9": "6",
	}
	rotated := make([]string, len(numSlice))
	for i := range numSlice {
		rotated = append(rotated, rotationMap[numSlice[i]])
	}

	return strings.Join(rotated, "") != numS
}

func isGoodV1(n int) bool {
	copyN := n
	notEq := false
	for copyN != 0 {
		dig := copyN % 10
		if dig == 3 || dig == 4 || dig == 7 {
			return false
		}

		if dig == 2 || dig == 5 || dig == 6 || dig == 9 {
			notEq = true
		}

		copyN /= 10
	}

	return notEq
}
