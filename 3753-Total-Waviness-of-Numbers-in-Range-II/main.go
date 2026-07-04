package main

import "fmt"

func totalWaviness(num1 int64, num2 int64) int64 {
	var count int64

	mem := map[string]struct{}{}

	for right := range 10 {
		for center := range 10 {
			for left := range 10 {
				if (right < center && left < center) || (right > center && left > center) {
					mem[fmt.Sprintf("%d%d%d", right, center, left)] = struct{}{}
				}
			}
		}
	}
	for i := num1; i <= num2; i++ {
		count += waves(mem, i)
	}
	return count
}

func waves(mem map[string]struct{}, num int64) int64 {
	var count int64
	for num/100 > 0 {
		if _, ok := mem[fmt.Sprintf("%03d", num%1000)]; ok {
			count++
		}
		num = num / 10
	}
	return count
}

func waveniess(num int) int {
	var count int
	for num/100 > 0 {
		right := num % 10
		center := (num % 100) / 10
		left := (num % 1000) / 100

		num = num / 10

		if (right < center && left < center) || (right > center && left > center) {
			count++
		}
	}

	return count
}
